package functions

/* WARNING: GENERATED CODE
 *
 * This file is generated by ow-specgen using metadata from src/github.com/flowmatters/openwater-core/models/functions/gate.go
 * 
 * Don't edit this file. Edit src/github.com/flowmatters/openwater-core/models/functions/gate.go instead!
 */
import (
//  "fmt"
  "github.com/flowmatters/openwater-core/sim"
  "github.com/flowmatters/openwater-core/data"
)


type Gate struct {
  
}

func (m *Gate) ApplyParameters(parameters data.ND2Float64) {

  
}


func buildGate() sim.TimeSteppingModel {
	result := Gate{}
	return &result
}

func init() {
	sim.Catalog["Gate"] = buildGate
}

func (m *Gate)  Description() sim.ModelDescription{
	var result sim.ModelDescription
	result.Parameters = []sim.ParameterDescription{
  }

  result.Inputs = []string{
  "trigger","incoming",}
  result.Outputs = []string{
  "outgoing",}

  result.States = []string{
  }

	return result
}




func (m *Gate) InitialiseStates(n int) data.ND2Float64 {
  // Zero states
	var result = data.NewArray2DFloat64(n,0)

	// for i := 0; i < n; i++ {
  //   stateSet := make(sim.StateSet,0)
  //   

  //   if result==nil {
  //     result = data.NewArray2DFloat64(stateSet.Len(0),n)
  //   }
  //   result.Apply([]int{0,i},[]int{1,1},stateSet)
	// }
 
	return result
}



func (m *Gate) Run(inputs data.ND3Float64, states data.ND2Float64, outputs data.ND3Float64) {

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
//	result.Outputs = data.NewArray3DFloat64( 1, inputLen, numCells)
//	result.States = states  //clone? make([]sim.StateSet, len(states))

  // fmt.Println("Running Gate for ",numCells,"cells")
  for i := 0; i < numCells; i++ {
    outputPosSlice[sim.DIMO_CELL] = i
    statesPosSlice[sim.DIMS_CELL] = i
    inputsPosSlice[sim.DIMI_CELL] = i%numInputSequences

    

    // fmt.Println("i",i)
    // fmt.Println("States",states.Shape())
    // fmt.Println("Tmp2",tmp2.Shape())
    

    
    
    

//    fmt.Println("is",inputDims,"tmpShape",tmpCI.Shape(),"cis",cellInputsShape)

		cellInputs := inputs.Slice(inputsPosSlice,inputsSizeSlice,nil).MustReshape(cellInputsShape)
//    fmt.Println("cellInputs Shape",cellInputs.Shape())
    
//    fmt.Println("{trigger <nil>}",tmpTS.Shape())
		trigger := cellInputs.Slice([]int{ 0,0}, []int{ 1,inputLen}, nil).MustReshape(inputNewShape).(data.ND1Float64)
    
//    fmt.Println("{incoming <nil>}",tmpTS.Shape())
		incoming := cellInputs.Slice([]int{ 1,0}, []int{ 1,inputLen}, nil).MustReshape(inputNewShape).(data.ND1Float64)
    

    

    
    
    outputPosSlice[sim.DIMO_OUTPUT] = 0
    outgoing := outputs.Slice(outputPosSlice,outputSizeSlice,outputStepSlice).MustReshape([]int{inputLen}).(data.ND1Float64)
    
    

		 gate(trigger,incoming,outgoing)

    
    
    

//		result.Outputs.ApplySpice([]int{i,0,0},[]int = make([]sim.Series, 1)
    
	}

//	return result
}
