package SendXml

import (
	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	"encoding/xml"
	"fmt"
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
		Postalcode	string		`xml:"postalcode"`
		City		string		`xml:"city"`
		Country		string		`xml:"country"`
	}
	
	firstName 	:= context.GetInput("firstname").(string)
	lastName  	:= context.GetInput("lastname").(string)
	streetname  	:= context.GetInput("streetname").(string)
	streetnumber	:= context.GetInput("streetnumber").(string)
	postalcode	:= context.GetInput("postalcode").(string)
	city	  	:= context.GetInput("city").(string)
	country   	:= context.GetInput("country").(string)
	
	transName 	:= lastName+", "+firstName
	transStreet	:= streetname+", "+streetnumber
	
	log.Infof("transName is [%s], and transStreet is [%s]",transName,transStreet)
	
	emp := &employee{Name: transName, Street: transStreet,Postalcode: postalcode,City: city, Country: country}
	emp.Comment = "Transformed XML" 
	
	output, err := xml.MarshalIndent(emp, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	
	context.SetOutput("XML", output)

	return true, nil
}
