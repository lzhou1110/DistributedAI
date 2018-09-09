package main
import (
	"fmt"
	"flag"
	"os"
	"io"
	"encoding/binary"
	gonn "gonn"
)
const pixelRange = 255
// 获取文件路径
func OpenFile(path string) *os.File {
	file, err := os.Open(path)
	if (err != nil) {
		fmt.Println(err)
		os.Exit(-1);
	}
	return file
}
// ReadMNISTImages 从路径获取数据, 放入labels
func ReadMNISTLabels(r io.Reader)(labels []byte) {
	header := [2]int32{}
	binary.Read(r, binary.BigEndian, &header)
	//fmt.Printf("%v\n", header[1])
	labels = make([]byte, header[1])
	r.Read(labels)
	return
}
// ReadMNISTImages 从路径获取放入Images
func ReadMNISTImages(r io.Reader)(images [][]byte, width, height int) {
	header := [4]int32{}
	binary.Read(r, binary.BigEndian, &header)
	images = make([][]byte, header[1])
	width, height = int(header[2]), int(header[3])
	for i := 0; i<len(images); i++ {

		images[i] = make([]byte, width * height)
		r.Read(images[i])
	}
	return
}
func pixel2Weight(px byte) float64 {
	return float64(px) / pixelRange *0.9 + 0.1;
}
// prepareX 获取imageData(60000, 784), 将其中的数据正则化
func prepareX(M [][]byte) [][]float64 {
	rows := len(M);
	result := make([][]float64, rows)
	for i := 0; i < rows; i++ {
		result[i] = make([]float64, len(M[i]));
		for j := 0; j < len(M[i]); j++ {
			result[i][j] = pixel2Weight(M[i][j]);
		}
	}
	return result;
}
// prepareY 接收到label[60000], 转为
func prepareY(N []byte) [][]float64 {
	result := make([][]float64, len(N));
	for i := 0; i < len(result); i++ {
		tmp := make([]float64, 10);
		for j := 0; j < 10; j++ {
			tmp[j] = 0.1;
		}
		tmp[N[i]] = 0.9;
		result[i] = tmp;
	}
	return result;
}
// 获取最大值
func argmax(A []float64) int {
	x := 0;
	v := -1.0;
	for i, a := range(A) {
		if a > v {
			x = i;
			v = a;
		}
	}
	return x;
}
func main() {
	//sourceLabelFile := flag.String("sl", "", "source label file")
	//source ImageFile := flag.String("si", "", "source image File")
	//testLabelFile := flag.String("tl", "", "test label file")
	//testImageFile := flag.String("ti", "", "test image file")
	sourceLabelFile := flag.String("sl", "src/train-labels-idx1-ubyte", "./train-labels-idx1-ubyte")
	sourceImageFile := flag.String("si", "src/train-images-idx3-ubyte", "./train-images-idx3-ubyte")
	testLabelFile := flag.String("tl", "src/t10k-labels-idx1-ubyte", "./t10k-labels-idx1-ubyte")
	testImageFile := flag.String("ti", "src/t10k-images-idx3-ubyte", "./t10k-images-idx3-ubyte")
	flag.Parse()

	if *sourceLabelFile == "" || *testLabelFile == ""{
		flag.Usage()
		os.Exit(-2)
	}
	fmt.Println("Loading training data...")
	// 获取label数据 image数据
	// go run mymnist.go -si train-images-idx3-ubyte -sl train-labels-idx1-ubyte  -ti t10k-images-idx3-ubyte   -tl t10k-labels-idx1-ubyte
	labelData := ReadMNISTLabels(OpenFile(*sourceLabelFile))  // go run example_mnist/mymnist.go didn't work
	fmt.Printf("%T %v\n",labelData, labelData[:10])
	imageData, width, height := ReadMNISTImages(OpenFile(*sourceImageFile))
	fmt.Println("imageData:", len(imageData),len(imageData[1]),width,height)
	fmt.Println("labelData:", len(labelData),labelData[0:10])
	//for i:=0; i<28; i++ {
	//	for j:=0; j<28; j++ {
	//		//fmt.Print(i*28+j);
	//		//fmt.Print(" ");
	//		fmt.Print(imageData[0][i*28+j]);
	//		fmt.Print(" ");
	//	}
	//	fmt.Println();
	//}
	inputs := prepareX(imageData); //(60000, 784)
	targets := prepareY(labelData); //(60000, 10)
	//fmt.Println(inputs[0]);
	//fmt.Println(targets[:1]);
	nn := gonn.NewNetwork(784, 100, 10, false, 0.25, 0.1);
	//fmt.Println(nn);
	nn.Train(inputs, targets, 10);

	// 测试数据
	var testLabelData []byte;
	var testImageData [][]byte;
	if *testLabelFile != "" && *testImageFile != "" {
		fmt.Println("Loading test data...");
		testLabelData = ReadMNISTLabels(OpenFile(*testLabelFile));
		testImageData, _, _ = ReadMNISTImages(OpenFile(*testImageFile));
	}

	test_inputs := prepareX(testImageData);
	test_targets := prepareY(testLabelData);
	//test_inputs = inputs[:1000]
	//test_targets = targets[:1000]

	correct_ct := 0;
	for i, p := range(test_inputs) {
		//fmt.Println(nn.Forward(p));
		y := argmax(nn.Forward(p));
		yy := argmax(test_targets[i]);
		if y == yy {
			correct_ct += 1;
		}
	}

	fmt.Println("correct rate: ", float64(correct_ct)/ float64(len(test_inputs)), correct_ct,len(test_inputs));
}
