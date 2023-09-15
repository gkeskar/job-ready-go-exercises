   var taxableIncome float64 = 140000
   var max10 float64 = 9875
   var max12 float64 = 40125
   var max22 float64 = 85525
   var max24 float64 = 163300
   var max32 float64 = 207350
   var max35 float64 = 518400

   var taxDue float64

   if taxableIncome <= max10 {
      taxDue = taxableIncome * 0.1
   }
   // etc