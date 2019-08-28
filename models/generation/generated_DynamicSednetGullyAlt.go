package generation

/* WARNING: GENERATED CODE
 *
 * This file is generated by ow-specgen using metadata from ./models/generation/sednet_gully_alt.go
 * 
 * Don't edit this file. Edit ./models/generation/sednet_gully_alt.go instead!
 */
import (
//  "fmt"
  "github.com/flowmatters/openwater-core/sim"
  "github.com/flowmatters/openwater-core/data"
)


type DynamicSednetGullyAlt struct {
  YearDisturbance data.ND1Float64
  GullyEndYear data.ND1Float64
  Area data.ND1Float64
  averageGullyActivityFactor data.ND1Float64
  AnnualRunoff data.ND1Float64
  GullyAnnualAverageSedimentSupply data.ND1Float64
  GullyPercentFine data.ND1Float64
  managementPracticeFactor data.ND1Float64
  annualLoad data.ND1Float64
  longtermRunoffFactor data.ND1Float64
  dailyRunoffPowerFactor data.ND1Float64
  sdrFine data.ND1Float64
  sdrCoarse data.ND1Float64
  
}

func (m *DynamicSednetGullyAlt) ApplyParameters(parameters data.ND2Float64) {

  nSets := parameters.Len(sim.DIMP_CELL)
  newShape := []int{nSets}

  m.YearDisturbance = parameters.Slice([]int{ 0, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.GullyEndYear = parameters.Slice([]int{ 1, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.Area = parameters.Slice([]int{ 2, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.averageGullyActivityFactor = parameters.Slice([]int{ 3, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.AnnualRunoff = parameters.Slice([]int{ 4, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.GullyAnnualAverageSedimentSupply = parameters.Slice([]int{ 5, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.GullyPercentFine = parameters.Slice([]int{ 6, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.managementPracticeFactor = parameters.Slice([]int{ 7, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.annualLoad = parameters.Slice([]int{ 8, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.longtermRunoffFactor = parameters.Slice([]int{ 9, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.dailyRunoffPowerFactor = parameters.Slice([]int{ 10, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.sdrFine = parameters.Slice([]int{ 11, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  m.sdrCoarse = parameters.Slice([]int{ 12, 0}, []int{ 1, nSets}, nil).MustReshape(newShape).(data.ND1Float64)
  
}


func buildDynamicSednetGullyAlt() sim.TimeSteppingModel {
	result := DynamicSednetGullyAlt{}
	return &result
}

func init() {
	sim.Catalog["DynamicSednetGullyAlt"] = buildDynamicSednetGullyAlt
}

func (m *DynamicSednetGullyAlt)  Description() sim.ModelDescription{
	var result sim.ModelDescription
	result.Parameters = []sim.ParameterDescription{
  
  sim.DescribeParameter("YearDisturbance",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("GullyEndYear",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("Area",0,"m^2",[]float64{ 0, 0 },""),
  sim.DescribeParameter("averageGullyActivityFactor",0,"",[]float64{ 0, 3 },""),
  sim.DescribeParameter("AnnualRunoff",0,"mm.yr^-1",[]float64{ 0, 0 },""),
  sim.DescribeParameter("GullyAnnualAverageSedimentSupply",0,"t.yr^-1",[]float64{ 0, 0 },""),
  sim.DescribeParameter("GullyPercentFine",0,"Average clay + silt percentage of gully material",[]float64{ 0, 0 },""),
  sim.DescribeParameter("managementPracticeFactor",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("annualLoad",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("longtermRunoffFactor",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("dailyRunoffPowerFactor",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("sdrFine",0,"",[]float64{ 0, 0 },""),
  sim.DescribeParameter("sdrCoarse",0,"",[]float64{ 0, 0 },""),}

  result.Inputs = []string{
  "quickflow","year",}
  result.Outputs = []string{
  "fineLoad","coarseLoad",}

  result.States = []string{
  }

	return result
}




func (m *DynamicSednetGullyAlt) InitialiseStates(n int) data.ND2Float64 {
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



func (m *DynamicSednetGullyAlt) Run(inputs data.ND3Float64, states data.ND2Float64, outputs data.ND3Float64) {

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
  // fmt.Println("Running DynamicSednetGullyAlt for ",numCells,"cells")
//  for i := 0; i < numCells; i++ {
  for j := 0; j < numCells; j++ {
    go func(i int){
      outputPosSlice := outputs.NewIndex(0)
      statesPosSlice := states.NewIndex(0)
      inputsPosSlice := inputs.NewIndex(0)

      outputPosSlice[sim.DIMO_CELL] = i
      statesPosSlice[sim.DIMS_CELL] = i
      inputsPosSlice[sim.DIMI_CELL] = i%numInputSequences

      
      // fmt.Println("YearDisturbance=",m.YearDisturbance)
      yeardisturbance := m.YearDisturbance.Get1(i%m.YearDisturbance.Len1())
      
      // fmt.Println("GullyEndYear=",m.GullyEndYear)
      gullyendyear := m.GullyEndYear.Get1(i%m.GullyEndYear.Len1())
      
      // fmt.Println("Area=",m.Area)
      area := m.Area.Get1(i%m.Area.Len1())
      
      // fmt.Println("averageGullyActivityFactor=",m.averageGullyActivityFactor)
      averagegullyactivityfactor := m.averageGullyActivityFactor.Get1(i%m.averageGullyActivityFactor.Len1())
      
      // fmt.Println("AnnualRunoff=",m.AnnualRunoff)
      annualrunoff := m.AnnualRunoff.Get1(i%m.AnnualRunoff.Len1())
      
      // fmt.Println("GullyAnnualAverageSedimentSupply=",m.GullyAnnualAverageSedimentSupply)
      gullyannualaveragesedimentsupply := m.GullyAnnualAverageSedimentSupply.Get1(i%m.GullyAnnualAverageSedimentSupply.Len1())
      
      // fmt.Println("GullyPercentFine=",m.GullyPercentFine)
      gullypercentfine := m.GullyPercentFine.Get1(i%m.GullyPercentFine.Len1())
      
      // fmt.Println("managementPracticeFactor=",m.managementPracticeFactor)
      managementpracticefactor := m.managementPracticeFactor.Get1(i%m.managementPracticeFactor.Len1())
      
      // fmt.Println("annualLoad=",m.annualLoad)
      annualload := m.annualLoad.Get1(i%m.annualLoad.Len1())
      
      // fmt.Println("longtermRunoffFactor=",m.longtermRunoffFactor)
      longtermrunofffactor := m.longtermRunoffFactor.Get1(i%m.longtermRunoffFactor.Len1())
      
      // fmt.Println("dailyRunoffPowerFactor=",m.dailyRunoffPowerFactor)
      dailyrunoffpowerfactor := m.dailyRunoffPowerFactor.Get1(i%m.dailyRunoffPowerFactor.Len1())
      
      // fmt.Println("sdrFine=",m.sdrFine)
      sdrfine := m.sdrFine.Get1(i%m.sdrFine.Len1())
      
      // fmt.Println("sdrCoarse=",m.sdrCoarse)
      sdrcoarse := m.sdrCoarse.Get1(i%m.sdrCoarse.Len1())
      

      // fmt.Println("i",i)
      // fmt.Println("States",states.Shape())
      // fmt.Println("Tmp2",tmp2.Shape())
      

      
      
      

  //    fmt.Println("is",inputDims,"tmpShape",tmpCI.Shape(),"cis",cellInputsShape)

      cellInputs := inputs.Slice(inputsPosSlice,inputsSizeSlice,nil).MustReshape(cellInputsShape)
  //    fmt.Println("cellInputs Shape",cellInputs.Shape())
      
  //    fmt.Println("{quickflow m^3.s^-1}",tmpTS.Shape())
      quickflow := cellInputs.Slice([]int{ 0,0}, []int{ 1,inputLen}, nil).MustReshape(inputNewShape).(data.ND1Float64)
      
  //    fmt.Println("{year year}",tmpTS.Shape())
      year := cellInputs.Slice([]int{ 1,0}, []int{ 1,inputLen}, nil).MustReshape(inputNewShape).(data.ND1Float64)
      

      

      
      
      outputPosSlice[sim.DIMO_OUTPUT] = 0
      fineload := outputs.Slice(outputPosSlice,outputSizeSlice,outputStepSlice).MustReshape([]int{inputLen}).(data.ND1Float64)
      
      outputPosSlice[sim.DIMO_OUTPUT] = 1
      coarseload := outputs.Slice(outputPosSlice,outputSizeSlice,outputStepSlice).MustReshape([]int{inputLen}).(data.ND1Float64)
      
      

       sednetGullyDerm(quickflow,year,yeardisturbance,gullyendyear,area,averagegullyactivityfactor,annualrunoff,gullyannualaveragesedimentsupply,gullypercentfine,managementpracticefactor,annualload,longtermrunofffactor,dailyrunoffpowerfactor,sdrfine,sdrcoarse,fineload,coarseload)

      
      
      

  //		result.Outputs.ApplySpice([]int{i,0,0},[]int = make([]sim.Series, 2)
      

      doneChan <- i
    }(j)
	}

  for j := 0; j < numCells; j++ {
    <- doneChan
  }
//	return result
}
