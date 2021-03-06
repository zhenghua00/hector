package hector

type Classifier interface {

	//Set training parameters from parameter map
	Init(params map[string]string)

	//Train model on a given dataset
	Train(dataset * DataSet)

	//Predict the probability of a sample to be positive sample
	Predict(sample * Sample) float64

	SaveModel(path string)
	LoadModel(path string)
}

type MultiClassClassifier interface {
	//Set training parameters from parameter map
	Init(params map[string]string)

	//Train model on a given dataset
	Train(dataset * DataSet)

	//Predict the probability of a sample to be positive sample
	PredictMultiClass(sample * Sample) *ArrayVector

	SaveModel(path string)
	LoadModel(path string)

}