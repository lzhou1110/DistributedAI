package gonn
import(
	"math/rand"
	"time"
)
type NeuralNetwork struct {
	HiddenLayer			[]float64
	InputLayer			[]float64
	OutputLayer			[]float64
	WeightHidden		[][]float64
	WeightOutput		[][]float64
	ErrOutput			[]float64
	ErrHidden			[]float64
	LastChangeHidden	[][]float64
	LastChangeOutput	[][]float64
	Regression			bool
	Rate1				float64	// learning rate
	Rate2				float64
}
// randomMatrix 初始化一个Matrix,赋值随机值
func randomMatrix(rows, colums int, lower, upper float64) [][]float64 {
	mat := make([][]float64, rows);
	for i:=0; i<rows; i++ {
		mat[i] = make([]float64, colums);
		for j:=0; j < colums; j++ {
			mat[i][j] = rand.Float64()*(upper-lower) + lower;
		}
	}
	return mat;
}

func makeMatrix(rows, colums int, value float64) [][]float64 {
	mat := make([][]float64, rows);
	for i:=0; i<rows; i++ {
		mat[i] = make([]float64, colums);
		for j:=0; j<colums; j++ {
			mat[i][j] = value;
		}
	}
	return mat;
}
// NewNeuralNetWork 新建一个神经网络
func NewNetwork(iInputCount, iHiddenCount, iOutputCount int, iRegression bool, iRate1, iRate2 float64) *NeuralNetwork {
	//fmt.Println();
	iInputCount += 1;
	iHiddenCount += 1;
	rand.Seed(time.Now().UnixNano());
	network := &NeuralNetwork{};
	network.Regression = iRegression;
	network.Rate1 = iRate1;
	network.Rate2 = iRate2;
	network.InputLayer = make([]float64, iInputCount);
	network.HiddenLayer = make([]float64, iHiddenCount);
	network.OutputLayer = make([]float64, iOutputCount);
	network.ErrOutput = make([]float64, iOutputCount);
	network.ErrHidden = make([]float64, iHiddenCount);

	network.WeightHidden = randomMatrix(iHiddenCount, iInputCount, -1.0, 1.0);
	network.WeightOutput = randomMatrix(iOutputCount, iHiddenCount, -1.0, 1.0);

	network.LastChangeHidden = makeMatrix(iHiddenCount, iInputCount, 0.0);
	network.LastChangeOutput = makeMatrix(iOutputCount, iHiddenCount, 0.0);

	return network;
}
// Train 训练数据
func (self *NeuralNetwork) Train(inputs [][]float64, targets [][]float64, iteration int) {
	if len(inputs[0]) + 1 != len(self.InputLayer) {
		panic("amount of input variable doesn't match");
	}
	if len(targets[0]) != len(self.OutputLayer) {
		panic("amount of output variable doesn't match");
	}
	iter_flag := -1;
	for
}