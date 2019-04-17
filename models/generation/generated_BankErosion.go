package generation

/* WARNING: GENERATED CODE
 *
 * This file is generated by ow-specgen using metadata from src/github.com/flowmatters/openwater-core/models/generation/bank_erosion.go
 * 
 * Don't edit this file. Edit src/github.com/flowmatters/openwater-core/models/generation/bank_erosion.go instead!
 */
import (
//  "fmt"
  "github.com/flowmatters/openwater-core/sim"
  "github.com/flowmatters/openwater-core/data"
)


type BankErosion struct {
  riparianVegPercent data.ND1Float64
  maxRiparianVegEffectiveness data.ND1Float64
  soilErodibility data.ND1Float64
  bankErosionCoeff data.ND1Float64
  linkSlope data.ND1Float64
  bankFullFlow data.ND1Float64
  bankMgtFactor data.ND1Float64
  sedBulkDensity data.ND1Float64
  bankHeight data.ND1Float64
  linkLength data.ND1Float64
  dailyFlowPowerFactor data.ND1Float64
  longTermAvDailyFlow data.ND1Float64
  soilPercentFine data.ND1Float64
  durationInSeconds data.ND1Float64
  
}

func (m *BankErosion) ApplyParameters(parameters data.ND2Float64) {

  nSets := parameters.Len(sim.DIMP_CELL)
  newShape := []int{nSets}

  m.riparianVegPercent = parameters.Slice([]int{ 0, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.maxRiparianVegEffectiveness = parameters.Slice([]int{ 1, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.soilErodibility = parameters.Slice([]int{ 2, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.bankErosionCoeff = parameters.Slice([]int{ 3, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.linkSlope = parameters.Slice([]int{ 4, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.bankFullFlow = parameters.Slice([]int{ 5, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.bankMgtFactor = parameters.Slice([]int{ 6, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.sedBulkDensity = parameters.Slice([]int{ 7, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.bankHeight = parameters.Slice([]int{ 8, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.linkLength = parameters.Slice([]int{ 9, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.dailyFlowPowerFactor = parameters.Slice([]int{ 10, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.longTermAvDailyFlow = parameters.Slice([]int{ 11, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.soilPercentFine = parameters.Slice([]int{ 12, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.durationInSeconds = parameters.Slice([]int{ 13, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  
}


func buildBankErosion() sim.TimeSteppingModel {
	result := BankErosion{}
	return &result
}

func init() {
	sim.Catalog["BankErosion"] = buildBankErosion
}

func (m *BankErosion)  Description() sim.ModelDescription{
	var result sim.ModelDescription
	result.Parameters = []sim.ParameterDescription{
  
  sim.DescribeParameter("riparianVegPercent",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("maxRiparianVegEffectiveness",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("soilErodibility",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("bankErosionCoeff",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("linkSlope",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("bankFullFlow",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("bankMgtFactor",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("sedBulkDensity",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("bankHeight",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("linkLength",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("dailyFlowPowerFactor",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("longTermAvDailyFlow",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("soilPercentFine",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("durationInSeconds",0,"",[]float64{ 0, 0 },""),}

  result.Inputs = []string{
  "downstreamFlowVolume","totalVolume",}
  result.Outputs = []string{
  "bankErosionFine","bankErosionCoarse",}

  result.States = []string{
  }

	return result
}




func (m *BankErosion) InitialiseStates(n int) data.ND2Float64 {
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



func (m *BankErosion) Run(inputs data.ND3Float64, states data.ND2Float64, outputs data.ND3Float64) {

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

//  outputPosSlice := outputs.NewIndex(0)
  outputStepSlice := outputs.NewIndex(1)
  outputSizeSlice := outputs.NewIndex(1)
  outputSizeSlice[sim.DIMO_TIMESTEP] = inputLen

//  statesPosSlice := states.NewIndex(0)
  statesSizeSlice := states.NewIndex(1)
  statesSizeSlice[sim.DIMS_STATE] = numStates

//  inputsPosSlice := inputs.NewIndex(0)
  inputsSizeSlice := inputs.NewIndex(1)
  inputsSizeSlice[sim.DIMI_INPUT] = inputDims[sim.DIMI_INPUT]
  inputsSizeSlice[sim.DIMI_TIMESTEP] = inputLen

//  var result sim.RunResults
//	result.Outputs = data.NewArray3DFloat64( 2, inputLen, numCells)
//	result.States = states  //clone? make([]sim.StateSet, len(states))

  doneChan := make(chan int)
  // fmt.Println("Running BankErosion for ",numCells,"cells")
//  for i := 0; i < numCells; i++ {
  for j := 0; j < numCells; j++ {
    go func(i int){
      outputPosSlice := outputs.NewIndex(0)
      statesPosSlice := states.NewIndex(0)
      inputsPosSlice := inputs.NewIndex(0)

      outputPosSlice[sim.DIMO_CELL] = i
      statesPosSlice[sim.DIMS_CELL] = i
      inputsPosSlice[sim.DIMI_CELL] = i%numInputSequences

      
      // fmt.Println("riparianVegPercent=",m.riparianVegPercent)
      riparianvegpercent := m.riparianVegPercent.Get1(i%m.riparianVegPercent.Len1())
      
      // fmt.Println("maxRiparianVegEffectiveness=",m.maxRiparianVegEffectiveness)
      maxriparianvegeffectiveness := m.maxRiparianVegEffectiveness.Get1(i%m.maxRiparianVegEffectiveness.Len1())
      
      // fmt.Println("soilErodibility=",m.soilErodibility)
      soilerodibility := m.soilErodibility.Get1(i%m.soilErodibility.Len1())
      
      // fmt.Println("bankErosionCoeff=",m.bankErosionCoeff)
      bankerosioncoeff := m.bankErosionCoeff.Get1(i%m.bankErosionCoeff.Len1())
      
      // fmt.Println("linkSlope=",m.linkSlope)
      linkslope := m.linkSlope.Get1(i%m.linkSlope.Len1())
      
      // fmt.Println("bankFullFlow=",m.bankFullFlow)
      bankfullflow := m.bankFullFlow.Get1(i%m.bankFullFlow.Len1())
      
      // fmt.Println("bankMgtFactor=",m.bankMgtFactor)
      bankmgtfactor := m.bankMgtFactor.Get1(i%m.bankMgtFactor.Len1())
      
      // fmt.Println("sedBulkDensity=",m.sedBulkDensity)
      sedbulkdensity := m.sedBulkDensity.Get1(i%m.sedBulkDensity.Len1())
      
      // fmt.Println("bankHeight=",m.bankHeight)
      bankheight := m.bankHeight.Get1(i%m.bankHeight.Len1())
      
      // fmt.Println("linkLength=",m.linkLength)
      linklength := m.linkLength.Get1(i%m.linkLength.Len1())
      
      // fmt.Println("dailyFlowPowerFactor=",m.dailyFlowPowerFactor)
      dailyflowpowerfactor := m.dailyFlowPowerFactor.Get1(i%m.dailyFlowPowerFactor.Len1())
      
      // fmt.Println("longTermAvDailyFlow=",m.longTermAvDailyFlow)
      longtermavdailyflow := m.longTermAvDailyFlow.Get1(i%m.longTermAvDailyFlow.Len1())
      
      // fmt.Println("soilPercentFine=",m.soilPercentFine)
      soilpercentfine := m.soilPercentFine.Get1(i%m.soilPercentFine.Len1())
      
      // fmt.Println("durationInSeconds=",m.durationInSeconds)
      durationinseconds := m.durationInSeconds.Get1(i%m.durationInSeconds.Len1())
      

      // fmt.Println("i",i)
      // fmt.Println("States",states.Shape())
      // fmt.Println("Tmp2",tmp2.Shape())
      

      
      
      

  //    fmt.Println("is",inputDims,"tmpShape",tmpCI.Shape(),"cis",cellInputsShape)

      cellInputs := inputs.Slice(inputsPosSlice,inputsSizeSlice,nil).MustReshape(cellInputsShape)
  //    fmt.Println("cellInputs Shape",cellInputs.Shape())
      
  //    fmt.Println("{downstreamFlowVolume <nil>}",tmpTS.Shape())
      downstreamflowvolume := cellInputs.Slice([]int{ 0,0}, []int{ 1,inputLen}, nil).MustReshape(inputNewShape).(data.ND1Float64)
      
  //    fmt.Println("{totalVolume <nil>}",tmpTS.Shape())
      totalvolume := cellInputs.Slice([]int{ 1,0}, []int{ 1,inputLen}, nil).MustReshape(inputNewShape).(data.ND1Float64)
      

      

      
      
      outputPosSlice[sim.DIMO_OUTPUT] = 0
      bankerosionfine := outputs.Slice(outputPosSlice,outputSizeSlice,outputStepSlice).MustReshape([]int{inputLen}).(data.ND1Float64)
      
      outputPosSlice[sim.DIMO_OUTPUT] = 1
      bankerosioncoarse := outputs.Slice(outputPosSlice,outputSizeSlice,outputStepSlice).MustReshape([]int{inputLen}).(data.ND1Float64)
      
      

       bankErosion(downstreamflowvolume,totalvolume,riparianvegpercent,maxriparianvegeffectiveness,soilerodibility,bankerosioncoeff,linkslope,bankfullflow,bankmgtfactor,sedbulkdensity,bankheight,linklength,dailyflowpowerfactor,longtermavdailyflow,soilpercentfine,durationinseconds,bankerosionfine,bankerosioncoarse)

      
      
      

  //		result.Outputs.ApplySpice([]int{i,0,0},[]int = make([]sim.Series, 2)
      

      doneChan <- i
    }(j)
	}

  for j := 0; j < numCells; j++ {
    <- doneChan
  }
//	return result
}