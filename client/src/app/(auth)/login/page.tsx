import SignUpForm from "@/components/auth/signup-form";

const SignUp = () => {
  return (
    <div className="flex gap-32">
      <div className="h-screen w-2/3 bg-white"></div>
      <SignUpForm />
    </div>
  );
};

export default SignUp;
