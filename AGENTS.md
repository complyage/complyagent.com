==== All Responses ===
success     : true/false
error       : string (message)
elapsed     : duration of call
data        : agent response

==== Agent Responses ====

agents/ocr - Model : PaddleOCR
      extract()
            request : 
                  image       : base 64 encoded image
            response
                  text        : string (extracted text

agents/face
      describe()
            request :
                  image       : base64 encoded image
            response 
                  age         : number
                  gender      : string
                  confidence  : 0.0 - 1.0
                  face        : base64 encoded image of the clipped face from the picture
      compare() 
            request : 
                  image1      : base 64 encoded image
                  image2      : base 64 encoded image
                  uuid        : unique facial recognition attribute
                  threshold   : number (how close to compare)

agents/nsfw - Model 
      describe()
            request 
                  image       : base64 encoded image
            response
                  nsfw        : boolean
                  confidence  : 0.0 - 1.0
                  classes     : [] of string i.e. [EXPOSED_BREAST]

agents/ollama - Model : Trained phi3
      prompt()
            request
                  prompt      : string
            response
                  response    : string of response

                  
