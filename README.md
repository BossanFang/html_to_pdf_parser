# html_to_pdf_parser

## Prerequisite:
Install wkhtmltopdf using below command 
  - ``sudo apt install wkhtmltopdf`` - for Ubuntu
  - ``brew install Caskroom/cask/wkhtmltopdf`` - For Mac

Get go package using the below command for parsing the HTMl template.
  - ``go get github.com/SebastiaanKlippert/go-wkhtmltopdf``

Get another go package for converting this html parse data to pdf using the below command.    
  - ``go get html/template``

## To run the project:   
Run below command 
  - ``go run main.go``