package main

import (
   //"encoding/json"
   //"io/ioutil"

   nn "neuralnetwork"
   "fmt"
   "flag"
   "os"
   "io"
   "encoding/binary"
)

type MNISTDataset struct {
   LearnSet [][]float64 `json:"learnset"`
   LearnLab []float64   `json:"learnlabel"`
   TestSet  [][]float64 `json:"testset"`
   TestLab  []float64   `json:"testlabel"`
}

// ------- data ---------
// ReadMNISTLabels读入label 文件数据，返回label数据make([]byte, header[1])
func ReadMNISTLabels (r io.Reader) (labels []byte) {
   header := [2]int32{}
   binary.Read(r, binary.BigEndian, &header)
   labels = make([]byte, header[1])
   r.Read(labels)
   return
}

// ReadMNISTImages 新建images shape为(60000, W*H)
func ReadMNISTImages (r io.Reader) (images [][]byte, width, height int) {
   header := [4]int32{}
   binary.Read(r, binary.BigEndian, &header)
   images = make([][]byte, header[1])  // [[] [] [] [] []]
   width, height = int(header[2]), int(header[3])
   for i := 0; i < len(images); i++ {
      images[i] = make([]byte, width * height)
      r.Read(images[i])
   }
   return
}
func OpenFile (path string) *os.File {
   file, err := os.Open(path)
   if (err != nil) {
      fmt.Println(err)
      os.Exit(-1)
   }
   return file
}
const pixelRange = 255
func pixelWeight (px byte) float64 {
   return float64(px) / pixelRange
}
// prepareX  接收到image(60000, W*H)数据,
func prepareX(M [][]byte) [][]float64{
   rows := len(M) // 60000
   result := make([][]float64,rows)
   for i:=0;i<rows;i++{
      result[i] = make([]float64,len(M[i]))
      for j:=0;j<len(M[i]);j++{
         result[i][j] = pixelWeight(M[i][j])
      }
   }
   return result
}

// prepareY 接收到labelData（60000)
func prepareY(N []byte) []float64{
   //result := make([][]float64,len(N))
   //for i:=0;i<len(result);i++{
   //   tmp := make([]float64,10)
   //   for j:=0;j<10;j++{
   //      tmp[j] = 0.1
   //   }
   //   tmp[N[i]] = 0.9
   //   result[i] = tmp
   //}
   result := make([]float64, len(N))
   for i:=0; i<len(result); i++ {
      result[i] = float64(N[i])
   }
   return result
}
// ------- data ---------

// ref: http://yann.lecun.com/exdb/mnist/
func MNISTDataLoad () []MNISTDataset{
   fmt.Println("Loading MNIST dataset ...")

   sourceLabelFile := flag.String("sl", "src/train-labels-idx1-ubyte", "source label file")
   sourceImageFile := flag.String("si", "src/train-images-idx3-ubyte", "source image file")
   testLabelFile := flag.String("tl", "src/t10k-labels-idx1-ubyte", "test label file")
   testImageFile := flag.String("ti", "src/t10k-images-idx3-ubyte", "test image file")
   flag.Parse()
   if *sourceLabelFile == "" || *sourceImageFile == "" {
      flag.Usage()
      os.Exit(-2)
   }
   fmt.Println("Loading training data...")
   labelData := ReadMNISTLabels(OpenFile(*sourceLabelFile))    // 获取label数据
   imageData, _, _ := ReadMNISTImages(OpenFile(*sourceImageFile))     // 获取image数据
   inputs := prepareX(imageData)
   targets := prepareY(labelData)
   //fmt.Println("inputs:", len(inputs),len(inputs[0]),width,height)
   //fmt.Println("targets:", len(targets),targets[0:10])
   var testLabelData []byte
   var testImageData [][]byte
   if *testLabelFile != "" && *testImageFile != "" {
      fmt.Println("Loading test data...")
      testLabelData = ReadMNISTLabels(OpenFile(*testLabelFile))
      testImageData, _, _ = ReadMNISTImages(OpenFile(*testImageFile))
   }
   test_inputs := prepareX(testImageData)
   test_targets := prepareY(testLabelData)
   //fmt.Println("test_inputs:", test_inputs[:10])
   //fmt.Println("test_targets:", test_targets[:10])
   // 转为json
   //dataSet := new(MNISTDataset)
   //dataSet.LearnSet = inputs
   //dataSet.LearnLab = targets
   //dataSet.TestSet = test_inputs
   //dataSet.TestLab = test_targets

   // 数据来源 按照index分类，0-9数组，再分别存入数据
   categorySet := make([]MNISTDataset, 10)
   fmt.Printf("categorySet: %T\n", categorySet[0])
   for i:=0; i<len(targets); i++ {
      label := int(targets[i])    // label-真实数字: 0 - 9
      //if i < 100 {
      //   fmt.Println(label)
      //}
      categorySet[label].LearnLab = append(categorySet[label].LearnLab, float64(label))
      categorySet[label].LearnSet = append(categorySet[label].LearnSet, inputs[i])
   }
   //fmt.Println("categorySet[label].LearnLab:", categorySet[1].LearnLab) // 理解意思啊。。
   //fmt.Println("categorySet[label].LearnSet:", categorySet[1].LearnSet[:2])
   for i:=0; i<len(test_targets); i++ {
      test_label := int(test_targets[i]);
      categorySet[test_label].TestSet = append(categorySet[test_label].TestSet, test_inputs[i])
      categorySet[test_label].TestLab = append(categorySet[test_label].TestLab, test_targets[i])
   }

   //var category_inputs0 [][]float64 = make([][]float64, 1000)
   //var category_targets0 []float64 = make([]float64, 1000)
   //var category_inputs1 [][]float64 = make([][]float64, 1000)
   //var category_targets1 []float64 = make([]float64, 1000)
   //var category_inputs2 [][]float64 = make([][]float64, 1000)
   //var category_targets2 []float64 = make([]float64, 1000)
   //
   //var category_test_inputs0 [][]float64 = make([][]float64, 1000)
   //var category_test_targets0 []float64 = make([]float64, 1000)
   //var category_test_inputs1 [][]float64 = make([][]float64, 1000)
   //var category_test_targets1 []float64 = make([]float64, 1000)
   //var category_test_inputs2 [][]float64 = make([][]float64, 1000)
   //var category_test_targets2 []float64 = make([]float64, 1000)
   //for i:=0; i <len(targets); i++ {
   //   index := int(targets[i])
   //   switch index {
   //   case 0:
   //      category_inputs0 = append(category_inputs0, inputs[index])
   //      category_targets0 = append(category_targets0, targets[i])
   //   case 1:
   //      category_inputs1 = append(category_inputs1, inputs[index])
   //      category_targets1 = append(category_targets1, targets[i])
   //   case 2:
   //      category_inputs2 = append(category_inputs2, inputs[index])
   //      category_targets2 = append(category_targets2, targets[i])
   //   }
   //}
   //
   //for i:=0; i <len(test_targets); i++ {
   //   index := int(test_targets[i])
   //   switch index {
   //   case 0:
   //      category_test_inputs0 = append(category_test_inputs0, inputs[index])
   //      category_test_targets0 = append(category_test_targets0, targets[i])
   //   case 1:
   //      category_test_inputs1 = append(category_test_inputs1, inputs[index])
   //      category_test_targets1 = append(category_test_targets1, targets[i])
   //   case 2:
   //      category_test_inputs2 = append(category_test_inputs2, inputs[index])
   //      category_test_targets2 = append(category_test_targets2, targets[i])
   //   }
   //}
   //
   //dataSet0 := new(MNISTDataset)
   //dataSet0.LearnSet = category_inputs0
   //dataSet0.LearnLab = category_targets0
   //dataSet0.TestSet = category_test_inputs0
   //dataSet0.TestLab = category_test_targets0
   //dataSet1 := new(MNISTDataset)
   //dataSet1.LearnSet = category_inputs1
   //dataSet1.LearnLab = category_targets1
   //dataSet1.TestSet = category_test_inputs1
   //dataSet1.TestLab = category_test_targets1
   //dataSet2 := new(MNISTDataset)
   //dataSet2.LearnSet = category_inputs2
   //dataSet2.LearnLab = category_targets2
   //dataSet2.TestSet = category_test_inputs2
   //dataSet2.TestLab = category_test_targets2
   //categorySet = append(categorySet, dataSet0)
   //categorySet = append(categorySet, dataSet1)
   //categorySet = append(categorySet, dataSet2)
   // 统一为小数
   //for _, image := range r.LearnSet {
   //   for i := len(image) - 1; i >= 0; i-- {
   //      image[i] *= 1.0 / 255.0
   //   }
   //}
   //for _, image := range r.TestSet {
   //   for i := len(image) - 1; i >= 0; i-- {
   //      image[i] *= 1.0 / 255.0
   //   }
   //}
   fmt.Println("Loaded.")
   return categorySet
}

func LabelEncodeVector (klass float64) *nn.SimpleMatrix {
   klass_vec := make([]float64, 10)
   klass_vec[int(klass)] = 1.0
   return nn.NewSimpleMatrix(1, 10).FillElt(klass_vec)
}

func LabelDecode (X *nn.SimpleMatrix) float64 {
   max := X.EltMax()
   for i, v := range X.Data[0] {
      if v == max {
         return float64(i)
      }
   }
   return -1
}

func LabelEqual (predict, expect *nn.SimpleMatrix) bool {
   return LabelDecode(predict) == LabelDecode(expect)
}

func __round__ (x float64) float64 {
   if x > 0.5 {
      return 1.0
   }
   return 0.0
}

// trainAndTest client k, 初始化参数 wt, batch_size, epoch, client k 所持数据
func trainAndTest(k int, wt *nn.SimpleMatrix, batch_size int, epoch int, dataset MNISTDataset) (*nn.SimpleMatrix, int) {
   // create CNN model
   nn.RandomSeed()
   n := nn.NewNeuralChain()
   n.AddLayer(nn.NewLayerConvolution(
      /* 1*1 image */ 1, 1, /* 12 filters */ 12,
      /* 28*28 pixels */ 28, 28,
      /* 5*5 kernel */ 5, 5,
      /* decay */ 0.001,
      ))
   n.AddLayer(nn.NewLayerActivation(1* 28, 12 * 28, "tanh"))
   n.AddLayer(nn.NewLayerPoolMax(1, 12, 28, 28, 2, 2))
   n.AddLayer(nn.NewLayerConvolution(1, 12, 16, 14, 14, 5, 5, 0.001))
   n.AddLayer(nn.NewLayerActivation(1 * 14, 16 * 14, "tanh"))
   n.AddLayer(nn.NewLayerFlatten(1 * 14, 16 * 14))
   n.AddLayer(nn.NewLayerLinear(1, 16 * 14 * 14, 10, 0.5, 0, true))
   n.AddLayer(nn.NewLayerLogRegression(1, 10))

   // data
   for index:=0; index<epoch; index++ {
      //trainSetNum := 60000
      //testSetNum := 10000
      trainSetNum := len(dataset.LearnSet)
      testSetNum := len(dataset.TestSet)
      error := 0
      for i := 0; i < trainSetNum; i++ {
         //k := rand.Intn(trainSetNum)
         image := nn.NewSimpleMatrix(28, 28).FillElt(dataset.LearnSet[i])
         label := LabelEncodeVector(dataset.LearnLab[i])
         predict := n.Predict(image)
         n.Learn(predict, label)
         n.Update(0.1)

         if !LabelEqual(predict, label) {
            error ++
         }
         if i % 1000 == 0 && i != 0{
            fmt.Printf("loss: %.2f%%\n", float64(error) / float64(i) * 100.0)
            //error = 0
            fmt.Println("Image", i, "->", dataset.LearnLab[i], label.Data[0], "  [A]", LabelDecode(predict), predict.Data[0])
         }
      }

      error = 0
      for i := 1; i <= testSetNum; i++ {
         k := i - 1
         image := nn.NewSimpleMatrix(28, 28).FillElt(dataset.TestSet[k])
         // TODO(fallenkliu@gamil.com): 正确率不对， 参数训练
         label := LabelEncodeVector(dataset.TestLab[k])  // 没有记录？
         predict := n.Predict(image)
         if !LabelEqual(predict, label) {
            error ++
         }
      }
      fmt.Printf("Test Loss: %.6f%%  correct: %.6f%%\n", float64(error) / float64(testSetNum), 1.0 - float64(error) / float64(testSetNum))
   }

   return wt, 1
}

// ------------- federating learning -------------
// clientUpdate client k client k-编号，传入的全局参数为wt； return updated weight and train nums
func clientUpdate(k int, wt *nn.SimpleMatrix, dataSet MNISTDataset) (*nn.SimpleMatrix, int){
   // 1. 数据集分块 根据Pk按batch_size为B进行分块
   B := 64
   batch_size := B
   epoch := 1
   fmt.Printf("client %d will begin training\n", k)
   // 2. 传入参数，训练数据(根据k进行识别编号)，测试数据，得到返回参数
   newWkt, nk := trainAndTest(k, wt, batch_size, epoch, dataSet)
   fmt.Println("当前参数newWkt:\n", newWkt)
   fmt.Println("当前client训练集个数：", nk)
   fmt.Printf("=================client %d train finished=================", k)
   return newWkt, nk
}

// server execute 参数: 全局的训练参数
func server(/*globalWeight nn.SimpleMatrix,*/ categorySet []MNISTDataset) {
   // 1. 初始化 全局参数
   var globalWeight *nn.SimpleMatrix;
   // 2. 总的全局训练次数
   t := 1
   wt := globalWeight   // 全局参数学习
   //sumN := 0         // 全局总的训练集个数
   for i:= 0; i < t; i++ {
      // 3. 选择一批client 集合
      // m←max(C ·K, 1): m 为c个clients和
      // St ←(random set of m clients) : 每次从m个中选取St个clients
      // st = np.array([0, 1, 2, 3, 4, 5, 6, 7, 8, 9])
      st := []int{0, 1, 2, 3}
      //var listAndnk []nn.SimpleMatrix  // 当前一次总的训练得到的参数

      // n 为所有clients 训练的训练数据集 当前循环t
      n := 0
      for _, k := range st {
         // newWkt 为当前client k的训练所得的参数, nk为client k的本地训练集个数
         newWkt, nk := clientUpdate(k, wt, categorySet[k]);
         n += nk
         //append(listAndnk, )
         fmt.Printf("index: %d \n", k)
         fmt.Println(newWkt)
      }
      // 4. 更新w权值(加权平均)
      // Wt+1 <-- sum(Nk/n * Wt+1): 即n个clients的参数加权求和得出全局的Wt+1

   }
}
// ------------- federating learning -------------

func main () {
   // operate dataSet
   categorySet := MNISTDataLoad()
   //mnist(MNISTDataLoad())
   server(categorySet)
}
