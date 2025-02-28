import { z } from "zod";

export const signUpSchema = z.object({
  username: z.string().min(3, "Username must be at least 3 characters"),
  name: z.string().min(3, "name must be at least 3 characters"),
  email: z.string().email("Please enter a valid email address"),
  password: z.string().min(8, "Password must be at least 8 characters"),
});

export const userSchema = signUpSchema.omit({ password: true });

export type UserSchema = z.infer<typeof userSchema>;
export type SignUpSchema = z.infer<typeof signUpSchema>;
