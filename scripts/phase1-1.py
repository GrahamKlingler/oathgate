from openai import OpenAI

c = OpenAI(base_url="http://localhost:8080/v1", api_key="unused")
r = c.chat.completions.create(model="llama3.2:1b", messages=[
    {
        "role": "user",
        "content": "Generate me an Israeli name"
}])
print(r.choices[0].message.content)