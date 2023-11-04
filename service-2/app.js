import express from "express"

const PORT = 3000

const app = express()

const handleGet = (req, res) => {
    res.json({
        message: "Hello from service-2"
    })
    return 
}

app.get("/", handleGet)

app.listen(PORT, () => {
    console.log(`App listens to on port ${PORT}`)
})
