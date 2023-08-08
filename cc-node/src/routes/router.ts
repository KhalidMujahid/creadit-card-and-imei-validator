import express, { NextFunction, Request, Response, Router } from "express";
import { creditCardValidator } from "../utils/validator";

const router: Router = express.Router();

router.get("/", (req: Request, res: Response, next: NextFunction) => {
  try {
    res.status(200).render("index");
  } catch (error) {
    next(error);
  }
});

router.get("/verify", (req: Request, res: Response, next: NextFunction) => {
  try {
    res.status(200).render("verify");
  } catch (error) {
    next(error);
  }
});

router.post(
  "/verify",
  (
    req: Request<{}, { value: string }, { value: string }>,
    res: Response<{ value: string }>,
    next: NextFunction
  ) => {
    const { value } = req.body;
    try {
      const result = creditCardValidator(value);
      if (result) {
        res.status(200).send({ value: "Credit card is valid" });
      } else {
        res.status(200).send({ value: "Credit card is not valid" });
      }
    } catch (error) {
      next(error);
    }
  }
);

export default router;
