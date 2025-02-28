import apiClient, { SuccessResponse } from "@/lib/api-client";
import { SignUpSchema, UserSchema } from "@/schema/auth/user";

export const signUpUser = async (user: SignUpSchema) => {
  return apiClient
    .post<SuccessResponse<Omit<UserSchema, "email">>>("api/auth/signup", {
      ...user,
    })
    .then((res) => res.data.data);
};
