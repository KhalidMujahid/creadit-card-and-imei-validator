import express, { Express } from "express";
import dotenv from "dotenv";
import * as middleware from "./middlewares/middleware";
import router from "./routes/router";

dotenv.config();

const app: Express = express();

app.set("view engine", "ejs");
app.use(express.static("public"));
app.use(express.json());

// routes
app.use("/", router);

// page not found handler
app.use(middleware.pageNotFound);

// error handler
app.use(middleware.errorHandler);

export default app;
