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
// 0.01 means sample 1% of MNIST dataset (train:600, test:100)
func MNISTDataLoad () *MNISTDataset{
   fmt.Println("Loading MNIST dataset ...")

   sourceLabelFile := flag.String("sl", "./train-labels-idx1-ubyte", "source label file")
   sourceImageFile := flag.String("si", "./train-images-idx3-ubyte", "source image file")
   testLabelFile := flag.String("tl", "./t10k-labels-idx1-ubyte", "test label file")
   testImageFile := flag.String("ti", "./t10k-images-idx3-ubyte", "test image file")
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
   dataSet := new(MNISTDataset)
   dataSet.LearnSet = inputs
   dataSet.LearnLab = targets
   dataSet.TestSet = test_inputs
   dataSet.TestLab = test_targets

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
   return dataSet
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

func mnist (dataset *MNISTDataset) {
   // create CNN model
   nn.RandomSeed()
   n := nn.NewNeuralChain()

   n.AddLayer(nn.NewLayerConvolution(
      /* 1*1 image */ 1, 1, /* 12 filters */ 12,
      /* 28*28 pixels */ 28, 28,
      /* 5*5 kernel */ 5, 5,
      /* decay */ 0.001))
   n.AddLayer(nn.NewLayerActivation(1* 28, 12 * 28, "tanh"))
   n.AddLayer(nn.NewLayerPoolMax(1, 12, 28, 28, 2, 2))
   n.AddLayer(nn.NewLayerConvolution(1, 12, 16, 14, 14, 5, 5, 0.001))
   n.AddLayer(nn.NewLayerActivation(1 * 14, 16 * 14, "tanh"))
   n.AddLayer(nn.NewLayerFlatten(1 * 14, 16 * 14))
   n.AddLayer(nn.NewLayerLinear(1, 16 * 14 * 14, 10, 0.5, 0, true))
   n.AddLayer(nn.NewLayerLogRegression(1, 10))

   // data
   trainSetNum := 60000
   testSetNum := 10000
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
      label := LabelEncodeVector(dataset.TestLab[k])
      predict := n.Predict(image)
      if !LabelEqual(predict, label) {
         error ++
      }
   }
   fmt.Printf("Test Loss: %.6f%%  correct: %.6f%%\n", float64(error) / float64(testSetNum), 1.0 - float64(error) / float64(testSetNum))
}


func main () {
   //MNISTDataLoad()
   mnist(MNISTDataLoad())
}
