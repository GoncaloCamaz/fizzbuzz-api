import express from "express";
import swaggerUi from "swagger-ui-express";
import fs from "fs";
import path from "path";
import { fileURLToPath } from "url";
import YAML from "yaml";

const __filename = fileURLToPath(import.meta.url);
const __dirname = path.dirname(__filename);

const app = express();
const PORT = process.env.PORT || 3000;

// load OpenAPI
const specPath = path.join(__dirname, "openapi.yaml");
const file = fs.readFileSync(specPath, "utf8");
const swaggerDocument = YAML.parse(file);

// serve raw yaml too
app.get("/openapi.yaml", (_req, res) => {
    res.type("text/yaml").send(file);
});

// swagger ui
app.use("/docs", swaggerUi.serve, swaggerUi.setup(swaggerDocument));

// health
app.get("/healthz", (_req, res) => res.json({ ok: true }));

app.listen(PORT, () => {
    console.log(`Swagger docs on http://localhost:${PORT}/docs`);
});