package main

func main() {
	r := registerRoutes()

	r.Run(":3000")

}

//////////////////////////////////////////////////////////
/////////////////////// Termianl ////////////////////////
////////////////////////////////////////////////////////
//                                                   //
// go mod init golang-gin-framework                 //
// go mod tidy                                     //
// cd src/app                                     //
// go mod init app                               //
// go mod tidy                                  //
// go build app                                //
// go run app                                 //
//                                           //
//////////////////////////////////////////////

//////////////////////////////////////////////////////////
//////////////////////// Browser ////////////////////////
////////////////////////////////////////////////////////
//                                                   //
// parameterized routes                             //
// http://localhost:3000/employees/962134/vacation //
//                                                //
// basic authentication                          //
// http://localhost:3000                        //
// user name: admin                            //
// password: admin                            //
//                                           //
//////////////////////////////////////////////
