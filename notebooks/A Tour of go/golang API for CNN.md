# 简介
接口文档针对neuralnetwork 模块。

# neural_chain 模块

1. 新建神经网络: `func NewNeuralChain () *NeuralChain`   
    `n := nn.NewNeuralChain()`
2. 添加网络层级: `func (n *NeuralChain) AddLayer (layer Layer) NeuralNetwork`
3. 新建convolution层: 
    ```
    func NewLayerConvolution (
   input_m, input_n, output_n int,
   item_m, item_n, kernel_m, kernel_n int,
   weight_decay float64,
) *LayerConvolution
    ```
4. 新建激活层: `func NewLayerActivation (input_m, input_n int, fun_type string) *LayerActivation`
5. 新建pooling层: `func NewLayerPoolMax (input_m, input_n, item_n, item_m, pool_m, pool_n int) *LayerPoolMax`
6. 新建最后的展开层: `func NewLayerFlatten (input_m, input_n int) *LayerFlatten`
7. 新建线性层: `func NewLayerLinear (input_m, input_n, output_n int, weight_scale, weight_decay float64, enable_b bool) *LayerLinear`
8. 新建分类层: `func NewLayerLogRegression (input_m, input_n int) *LayerLogRegression`
9. 新建矩阵: `func NewSimpleMatrix (m, n int) *SimpleMatrix `
10. 新建训练算法: `func (n *NeuralChain) Predict (input *SimpleMatrix) *SimpleMatrix`
11. Learn算法: `func (n *NeuralChain) Learn (predict *SimpleMatrix, expect *SimpleMatrix) NeuralNetwork`
12. 更新算法: `func (n *NeuralChain) Update (alpha float64) NeuralNetwork`


