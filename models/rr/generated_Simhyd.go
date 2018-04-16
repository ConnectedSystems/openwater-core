package rr

/* WARNING: GENERATED CODE
 *
 * This file is generated by ow-specgen using metadata from src/github.com/flowmatters/openwater-core/models/rr/simhyd.go
 * 
 * Don't edit this file. Edit src/github.com/flowmatters/openwater-core/models/rr/simhyd.go instead!
 */
import (
//  "fmt"
  "github.com/flowmatters/openwater-core/sim"
  "github.com/flowmatters/openwater-core/data"
)


type Simhyd struct {
  baseflowCoefficient data.ND1Float64
  imperviousThreshold data.ND1Float64
  infiltrationCoefficient data.ND1Float64
  infiltrationShape data.ND1Float64
  interflowCoefficient data.ND1Float64
  perviousFraction data.ND1Float64
  rainfallInterceptionStoreCapacity data.ND1Float64
  rechargeCoefficient data.ND1Float64
  soilMoistureStoreCapacity data.ND1Float64
  
}

func (m *Simhyd) ApplyParameters(parameters data.ND2Float64) {
  // fmt.Println(parameters)
  // fmt.Println(parameters.Shape())
  nSets := parameters.Len(sim.DIMP_CELL)
  // fmt.Println(nSets)
  m.baseflowCoefficient = parameters.Slice([]int{ 0, 0}, []int{ 1, nSets}, nil).(data.ND1Float64)
  m.imperviousThreshold = parameters.Slice([]int{ 1, 0}, []int{ 1, nSets}, nil).(data.ND1Float64)
  m.infiltrationCoefficient = parameters.Slice([]int{ 2, 0}, []int{ 1, nSets}, nil).(data.ND1Float64)
  m.infiltrationShape = parameters.Slice([]int{ 3, 0}, []int{ 1, nSets}, nil).(data.ND1Float64)
  m.interflowCoefficient = parameters.Slice([]int{ 4, 0}, []int{ 1, nSets}, nil).(data.ND1Float64)
  m.perviousFraction = parameters.Slice([]int{ 5, 0}, []int{ 1, nSets}, nil).(data.ND1Float64)
  m.rainfallInterceptionStoreCapacity = parameters.Slice([]int{ 6, 0}, []int{ 1, nSets}, nil).(data.ND1Float64)
  m.rechargeCoefficient = parameters.Slice([]int{ 7, 0}, []int{ 1, nSets}, nil).(data.ND1Float64)
  m.soilMoistureStoreCapacity = parameters.Slice([]int{ 8, 0}, []int{ 1, nSets}, nil).(data.ND1Float64)
  
}


func buildSimhyd() sim.TimeSteppingModel {
	result := Simhyd{}
	return &result
}

func init() {
	sim.Catalog["Simhyd"] = buildSimhyd
}

func (m *Simhyd)  Description() sim.ModelDescription{
	var result sim.ModelDescription
	result.Parameters = []sim.ParameterDescription{
  
  sim.DescribeParameter("baseflowCoefficient",0,""),
  sim.DescribeParameter("imperviousThreshold",0,""),
  sim.DescribeParameter("infiltrationCoefficient",0,""),
  sim.DescribeParameter("infiltrationShape",0,""),
  sim.DescribeParameter("interflowCoefficient",0,""),
  sim.DescribeParameter("perviousFraction",0,""),
  sim.DescribeParameter("rainfallInterceptionStoreCapacity",0,""),
  sim.DescribeParameter("rechargeCoefficient",0,""),
  sim.DescribeParameter("soilMoistureStoreCapacity",0,""),}

  result.Inputs = []string{
  "rainfall","pet",}
  result.Outputs = []string{
  "runoff","baseflow","store",}

  result.States = []string{
  "SoilMoistureStore","Groundwater","TotalStore",}

	return result
}




func (m *Simhyd) InitialiseStates(n int) data.ND2Float64 {
  // Zero states
	var result = data.NewArray2DFloat64(n,3)

	// for i := 0; i < n; i++ {
  //   stateSet := make(sim.StateSet,3)
  //   
	// 	stateSet[0] = 0 // SoilMoistureStore
  //   
	// 	stateSet[1] = 0 // Groundwater
  //   
	// 	stateSet[2] = 0 // TotalStore
  //   

  //   if result==nil {
  //     result = data.NewArray2DFloat64(stateSet.Len(0),n)
  //   }
  //   result.Apply([]int{0,i},[]int{1,1},stateSet)
	// }
 
	return result
}



func (m *Simhyd) Run(inputs data.ND3Float64, states data.ND2Float64, outputs data.ND3Float64) {

  // Loop over all cells
  inputDims := inputs.Shape()
  numCells := states.Len(sim.DIMS_CELL)
  numStates := states.Len(sim.DIMS_STATE)
  numInputSequences := inputs.Len(sim.DIMI_CELL)

  //  fmt.Println("num cells",lenStates,"num states",numStates)
  // fmt.Println("states shape",states.Shape())
  // fmt.Println("states",states) 
  inputLen := inputDims[sim.DIMI_TIMESTEP]
  cellInputsShape := inputDims[1:]
  inputNewShape := []int{inputLen}

  outputPosSlice := outputs.NewIndex(0)
  outputStepSlice := outputs.NewIndex(1)
  outputSizeSlice := outputs.NewIndex(1)
  outputSizeSlice[sim.DIMO_TIMESTEP] = inputLen

  statesPosSlice := states.NewIndex(0)
  statesSizeSlice := states.NewIndex(1)
  statesSizeSlice[sim.DIMS_STATE] = numStates

  inputsPosSlice := inputs.NewIndex(0)
  inputsSizeSlice := inputs.NewIndex(1)
  inputsSizeSlice[sim.DIMI_INPUT] = inputDims[sim.DIMI_INPUT]
  inputsSizeSlice[sim.DIMI_TIMESTEP] = inputLen

//  var result sim.RunResults
//	result.Outputs = data.NewArray3DFloat64( 4, inputLen, numCells)
//	result.States = states  //clone? make([]sim.StateSet, len(states))

  // fmt.Println("Running Simhyd for ",numCells,"cells")
  for i := 0; i < numCells; i++ {
    outputPosSlice[sim.DIMO_CELL] = i
    statesPosSlice[sim.DIMS_CELL] = i
    inputsPosSlice[sim.DIMI_CELL] = i%numInputSequences

    
    // fmt.Println("baseflowCoefficient=",m.baseflowCoefficient)
		baseflowcoefficient := m.baseflowCoefficient.Get1(i%m.baseflowCoefficient.Len1())
    
    // fmt.Println("imperviousThreshold=",m.imperviousThreshold)
		imperviousthreshold := m.imperviousThreshold.Get1(i%m.imperviousThreshold.Len1())
    
    // fmt.Println("infiltrationCoefficient=",m.infiltrationCoefficient)
		infiltrationcoefficient := m.infiltrationCoefficient.Get1(i%m.infiltrationCoefficient.Len1())
    
    // fmt.Println("infiltrationShape=",m.infiltrationShape)
		infiltrationshape := m.infiltrationShape.Get1(i%m.infiltrationShape.Len1())
    
    // fmt.Println("interflowCoefficient=",m.interflowCoefficient)
		interflowcoefficient := m.interflowCoefficient.Get1(i%m.interflowCoefficient.Len1())
    
    // fmt.Println("perviousFraction=",m.perviousFraction)
		perviousfraction := m.perviousFraction.Get1(i%m.perviousFraction.Len1())
    
    // fmt.Println("rainfallInterceptionStoreCapacity=",m.rainfallInterceptionStoreCapacity)
		rainfallinterceptionstorecapacity := m.rainfallInterceptionStoreCapacity.Get1(i%m.rainfallInterceptionStoreCapacity.Len1())
    
    // fmt.Println("rechargeCoefficient=",m.rechargeCoefficient)
		rechargecoefficient := m.rechargeCoefficient.Get1(i%m.rechargeCoefficient.Len1())
    
    // fmt.Println("soilMoistureStoreCapacity=",m.soilMoistureStoreCapacity)
		soilmoisturestorecapacity := m.soilMoistureStoreCapacity.Get1(i%m.soilMoistureStoreCapacity.Len1())
    

    // fmt.Println("i",i)
    // fmt.Println("States",states.Shape())
    // fmt.Println("Tmp2",tmp2.Shape())
    
    initialStates := states.Slice(statesPosSlice,statesSizeSlice,nil).MustReshape([]int{numStates}).(data.ND1Float64)
    
    // fmt.Println("IS Shape",initialStates.Shape())
    // fmt.Println("IS",initialStates)

    
    
    soilmoisturestore := initialStates.Get1(0)
    
    groundwater := initialStates.Get1(1)
    
    totalstore := initialStates.Get1(2)
    
    

//    fmt.Println("is",inputDims,"tmpShape",tmpCI.Shape(),"cis",cellInputsShape)

		cellInputs := inputs.Slice(inputsPosSlice,inputsSizeSlice,nil).MustReshape(cellInputsShape)
//    fmt.Println("cellInputs Shape",cellInputs.Shape())
    
//    fmt.Println("{rainfall mm}",tmpTS.Shape())
		rainfall := cellInputs.Slice([]int{ 0,0}, []int{ 1,inputLen}, nil).MustReshape(inputNewShape).(data.ND1Float64)
    
//    fmt.Println("{pet mm}",tmpTS.Shape())
		pet := cellInputs.Slice([]int{ 1,0}, []int{ 1,inputLen}, nil).MustReshape(inputNewShape).(data.ND1Float64)
    

    

    
    
    outputPosSlice[sim.DIMO_OUTPUT] = 0
    runoff := outputs.Slice(outputPosSlice,outputSizeSlice,outputStepSlice).MustReshape([]int{inputLen}).(data.ND1Float64)
    
    outputPosSlice[sim.DIMO_OUTPUT] = 1
    baseflow := outputs.Slice(outputPosSlice,outputSizeSlice,outputStepSlice).MustReshape([]int{inputLen}).(data.ND1Float64)
    
    outputPosSlice[sim.DIMO_OUTPUT] = 2
    store := outputs.Slice(outputPosSlice,outputSizeSlice,outputStepSlice).MustReshape([]int{inputLen}).(data.ND1Float64)
    
    

		soilmoisturestore,groundwater,totalstore= simhyd(rainfall,pet,soilmoisturestore,groundwater,totalstore,baseflowcoefficient,imperviousthreshold,infiltrationcoefficient,infiltrationshape,interflowcoefficient,perviousfraction,rainfallinterceptionstorecapacity,rechargecoefficient,soilmoisturestorecapacity,runoff,baseflow,store)

    
    
    initialStates.Set1(0, soilmoisturestore)
    
    initialStates.Set1(1, groundwater)
    
    initialStates.Set1(2, totalstore)
    
    

//		result.Outputs.ApplySpice([]int{i,0,0},[]int = make([]sim.Series, 3)
    
	}

//	return result
}
