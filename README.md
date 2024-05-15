## Simple math program to evaluate expressions
### Setup
run `make build` - to build an run the docker image

run `make run` - to just run the docker image without rebuild

run `make build_local` - to compile the `api` and `cli` inside `build/`
### Usage
There are 3 endpoint: `/evaluate`, `/validate`, `/errors`
#### 1. Making requests manually with curl
##### POST `/evaluate`
```
curl -X POST -H "Content-Type: application/json" -d '{"expression":"what is 2 plus 3?"}' localhost:55555/evaluate         
```
`{"result":5}%`
```
curl -X POST -H "Content-Type: application/json" -d '{"expression":"what is 2 divided by 0?"}' localhost:55555/evaluate               
```
`{"error":"division by zero"}%`

##### POST `/validate`
```
curl -X POST -H "Content-Type: application/json" -d '{"expression":"what is 3 multiplied by 3?"}' localhost:55555/validate
```            
`{"valid":true}%`
```
curl -X POST -H "Content-Type: application/json" -d '{"expression":"what is 3 multiplied by 3 cubed 2?"}' localhost:55555/validate    
```
`{"valid":false,"reason":"Unsupported error: Unsupported operation \"cubed\" at 26-31"}%`
##### GET `/errors`
```
curl -X GET localhost:55555/errors                                                                                                    
```
`[{"endpoint":"/evaluate","expression":"what is 2 squared?","frequency":"1","type":"Unsupported operation error"},{"endpoint":"/validate","
expression":"what is 3 multiplied by 3 cubed 2?","frequency":"1","type":"Unsupported operation error"}]%`
#### 2. Making request with the cli client
Run `./cli -m [eval|validate|errors]` This will put you in a interactive prompt where you directly type expressions and recieve results

`/.cli`
```
✦ ❯ ./cli                                                                                                                                 
Enter text (Ctrl+D to end):                                                                                                               
                                                                                                                                          
evaluate > What is 2?                                                                                                                     
{                                                                                                                                         
  "result": 2                                                                                                                             
}                                                                                                                                         
                                                                                                                                          
evaluate > what is 2 plus 3?                                                                                                              
{                                                                                                                                         
  "result": 5                                                                                                                             
}                                                                                                                                         
                                                                                                                                          
evaluate > What is 1 plus?                                                                                                                
{                                                                                                                                         
  "error": "Syntax error: expected number, got \"?\" at 14-15"                                                                            
}                                                                                                                                         
                                                                                                                                          
evaluate >    
```
`./cli -m validate`
```
✦ ❯ ./cli -m validate                                                                                                                     
Enter text (Ctrl+D to end):                                                                                                               
                                                                                                                                          
validate > What?                                                                                                                          
{                                                                                                                                         
  "valid": false,                                                                                                                         
  "reason": "non-math question(no numbers found in question)"                                                                             
}                                                                                                                                         
                                                                                                                                          
validate > What is 2                                                                                                                      
{                                                                                                                                         
  "valid": false,                                                                                                                         
  "reason": "Syntax error: expected termination keyword \"?\", after \"2\" at 8-9"                                                        
}                                                                                                                                         
                                                                                                                                          
validate > What is 6 divided by 2?                                                                                                        
{                                                                                                                                         
  "valid": true                                                                                                                           
}                                                                                                                                         
                                                                                                                                          
validate >
```
`./cli -m errors` This will directly print all the accumulated errors
```
✦ ❯ ./cli -m errors                                                                                                                       
[                                                                                                                                         
  {                                                                                                                                       
    "endpoint": "/validate",                                                                                                              
    "expression": "what is 3 multiplied by 3 cubed 2?",                                                                                   
    "frequency": "1",                                                                                                                     
    "type": "Unsupported operation error"                                                                                                 
  },                                                                                                                                      
  {                                                                                                                                       
    "endpoint": "/evaluate",                                                                                                              
    "expression": "What is 1 plus?",                                                                                                      
    "frequency": "1",                                                                                                                     
    "type": "Syntax error"                                                                                                                
  },                                                                                                                                      
  {                                                                                                                                       
    "endpoint": "/evaluate",                                                                                                              
    "expression": "./cli validate",                                                                                                       
    "frequency": "1",                                                                                                                     
    "type": "Non-math question error"                                                                                                     
  },                                                                                                                                      
  {                                                                                                                                       
    "endpoint": "/validate",                                                                                                              
    "expression": "What?",                                                                                                                
    "frequency": "1",                                                                                                                     
    "type": "Non-math question error"                                                                                                     
  },                                                                                                                                      
  {                                                                                                                                       
    "endpoint": "/validate",                                                                                                              
    "expression": "What is 2",                                                                                                            
    "frequency": "1",                                                                                                                     
    "type": "Syntax error"                                                                                                                
  },                                                                                                                                      
  {                                                                                                                                       
    "endpoint": "/evaluate",                                                                                                              
    "expression": "what is 2 squared?",                                                                                                   
    "frequency": "1",                                                                                                                     
    "type": "Unsupported operation error"                                                                                                 
  }                                                                                                                                       
]
```
