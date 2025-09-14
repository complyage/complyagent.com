# ComplyAgent :: Agents Documentation

## 📦 All Responses

Every agent response follows a common schema:

```json
{
  "success": true/false,
  "error": "string (message)",
  "elapsed": "duration of call",
  "data": { "agent specific response" }
}
```

---

## 🤖 Agent Responses

### 1. OCR Agent
**Model:** PaddleOCR

**Method: `extract()`**  
- **Request:**  
  - `image` : base64 encoded image  

- **Response:**  
  - `text` : string (extracted text)

---

### 2. Face Agent
**Methods:**

#### `describe()`  
- **Request:**  
  - `image` : base64 encoded image  

- **Response:**  
  - `age` : number  
  - `gender` : string  
  - `confidence` : 0.0 - 1.0  
  - `face` : base64 encoded image of the clipped face from the picture  

#### `compare()`  
- **Request:**  
  - `image1` : base64 encoded image  
  - `image2` : base64 encoded image  
  - `uuid` : unique facial recognition attribute  
  - `threshold` : number (similarity threshold)  

- **Response:**  
  - `match` : boolean  
  - `confidence` : 0.0 - 1.0  

---

### 3. NSFW Agent
**Model:** NudeNet Detector  

**Method: `describe()`**  
- **Request:**  
  - `image` : base64 encoded image  

- **Response:**  
  - `nsfw` : boolean  
  - `confidence` : 0.0 - 1.0  
  - `classes` : array of strings, e.g. `[ "EXPOSED_BREAST" ]`  

---

### 4. Ollama Agent
**Model:** Trained Phi-3  

**Method: `prompt()`**  
- **Request:**  
  - `prompt` : string  

- **Response:**  
  - `response` : string of model output  
