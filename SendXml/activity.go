package SendXml

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"encoding/xml"
	"fmt"
	"os"
)

//Logger
	var log = logger.GetLogger("activity-SendXml")

// MyActivity is a stub for your Activity implementation
type MyActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new activity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &MyActivity{metadata: metadata}
}

// Metadata implements activity.Activity.Metadata
func (a *MyActivity) Metadata() *activity.Metadata {
	return a.metadata
}

// Eval implements activity.Activity.Eval
func (a *MyActivity) Eval(context activity.Context) (done bool, err error)  {

	// do eval
	
	type employee struct{
		XMLName		xml.Name	`xml:"employee"`
		Comment 	string 		`xml:",comment"`
		Name		string		`xml:"name"`
		Street		string		`xml:"street"`
		Postalcode	int64		`xml:"postalcode"`
		City		string		`xml:"city"`
		Country		string		`xml:"country"`
	}
	
	firstName 	:= context.GetInput("firstname").(string)
	lastName  	:= context.GetInput("lastname").(string)
	streetname  	:= context.GetInput("streetname").(string)
	streetnumber	:= context.GetInput("streetnumber").(int64)
	postalcode	:= context.GetInput("postalcode").(int64)
	city	  	:= context.GetInput("city").(string)
	country   	:= context.GetInput("country").(string)
	
	transName 	:= lastName + firstName
	transStreet	:= streetname + string(streetnumber)
	
	log.Infof("transName is [%s], and transStreet is [%s]",transName,transStreet)
	
	emp := &employee{Name: transName, Street: transStreet,Postalcode: postalcode,City: city, Country: country}
	emp.Comment = "Transformed XML" 
	
	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(emp); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	//log.Infof("String Output is [%s]",s)
	log.Infof("enc is [%s]",enc)
	
	context.SetOutput("XML", enc)

	return true, nil
}
