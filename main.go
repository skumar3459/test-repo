package main
//import "github.com/aws/aws-lambda-go/lambda"

func handleRequest () (string, error) {
    return "Hello from Go!!!!!!!!", nil
}
func main() {
    handleRequest()
   // lambda.Start(handleRequest)
}