"use client";

import { useState, ChangeEvent, FormEvent } from "react";
import { Button } from "../ui/button";
import { Input } from "../ui/input";
import { z } from "zod";
import { signUpSchema, SignUpSchema } from "@/schema/auth/user";

type FormErrors = {
  [K in keyof SignUpSchema]?: string;
};

const SignUpForm = () => {
  const [formData, setFormData] = useState<SignUpSchema>({
    username: "",
    email: "",
    password: "",
  });

  const [errors, setErrors] = useState<FormErrors>({});

  const handleChange = (e: ChangeEvent<HTMLInputElement>): void => {
    const { name, value } = e.target;
    setFormData({
      ...formData,
      [name]: value,
    });

    if (errors[name as keyof SignUpSchema]) {
      setErrors({
        ...errors,
        [name]: undefined,
      });
    }
  };

  const handleSubmit = (e: FormEvent<HTMLFormElement>): void => {
    e.preventDefault();

    try {
      signUpSchema.parse(formData);

      console.log("Form submitted:", formData);
      alert("Sign up successful!");
    } catch (error) {
      if (error instanceof z.ZodError) {
        const newErrors: FormErrors = {};
        error.errors.forEach((err) => {
          const path = err.path[0] as keyof SignUpSchema;
          newErrors[path] = err.message;
        });
        setErrors(newErrors);
      }
    }
  };

  return (
    <form onSubmit={handleSubmit} className="mt-4 space-y-2">
      <div>
        <Input
          className="h-[45px] w-full border-none bg-neutral-900 px-4 leading-none text-white"
          placeholder="Enter your username"
          name="username"
          value={formData.username}
          onChange={handleChange}
        />
        {errors.username && (
          <p className="mt-1 text-sm text-red-500">{errors.username}</p>
        )}
      </div>

      <div>
        <Input
          className="h-[45px] w-full border-none bg-neutral-900 px-4 leading-none text-white"
          placeholder="Enter your email"
          type="email"
          name="email"
          value={formData.email}
          onChange={handleChange}
        />
        {errors.email && (
          <p className="mt-1 text-sm text-red-500">{errors.email}</p>
        )}
      </div>

      <div>
        <Input
          className="h-[45px] w-full border-none bg-neutral-900 px-4 leading-none text-white"
          placeholder="Enter your password"
          type="password"
          name="password"
          value={formData.password}
          onChange={handleChange}
        />
        {errors.password && (
          <p className="mt-1 text-sm text-red-500">{errors.password}</p>
        )}
      </div>

      <Button type="submit" className="mt-4 w-full cursor-pointer">
        Sign up
      </Button>
    </form>
  );
};

export default SignUpForm;
