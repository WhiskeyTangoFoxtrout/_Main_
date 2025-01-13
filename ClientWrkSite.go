/*I need some sites so I can at least look at my servers work
<!--How can I get you to post something and have it on yo orgin page
ima basically make the home feed a randomized page of basically dms and videos you like*/
//this bcam a go file onna go

import(
    "html/template"
    "os"
    "net/http"
)

<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE-Edge">
    <meta name="viewport" content="width=device-width, initial-scaled=1.0">
    <meta name="description" content="AfterLognCLientInitView">
    <meta name="keywords" content="HTML,CSS,GO">
    <title>Feeder Chamber</title>
</head>
<body>
    <h2>You should be able to read this</h2>
    <div class="postButton" aref="apawPic">
        <input type="button" href="#paw"><a href="#PostBox">paws</a>/*If you click it a box should show up*/
        <div class="print">
            <a class="paw1"></a><!--This is setting-->
            <a class="paw2"></a><!--This is to the search page-->
            <a class="paw3"></a><!--Just a feed page-->
            <a class="paw4"></a><!---->
        </div>
    </div>
    <div class="SearchBar">
        <a class="active" href="#home"></a>Home<!--# is a name anchor-->
        <a href="#AccountOptions">Account Options</a>
        <a href="#PageDets" >Page Details</a>
        <input type="text" placeholder="Search for NetSites">
    </div>
    <div class="PostBox"><!--ehat happens when you click it-->
        <input type="text" placeholder="Whats on your mind">
    </div>
</body>
//crap non of this should be http

func handler(){

    nil //getting somewhere
    
}

func main(){
    nil
    //gottta hook this up to the IPv6 on the clientside and I need a handlerFunc
}
//ima have to get rid of all this http
func renderTemplate(w http.ResponseWriter, r *http.Resquest){

    CliWrkSite := r.URL.Path[len("/view/"):]
    p,_:= loadPage(CliWrkSite)
    renderTemplate(w, "view", p)
    
}

func viewHandler(w http.ResponseWriter, r *http.Request){//hoping this hook up to the CliHoServe bc it got the URL

    CliWrkSite := r.URL.Path[len("/view/"):]
    p,_ := loadPage(CliWrkSite)
    renderTemplate(w, "view", p)
    
}

func editHandler(w http.ResponseWriter, r *http.Request){

    CliWrkSite := r.URL.Path[len("/view/"):]
    p, err := loadPage(CliWrkSite)
    if err != nil{

        p = &Page{CliWrkSite:CliWrkSite}
        
    }
    renderTemplate
}
