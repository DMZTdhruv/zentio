import SignUpForm from "@/components/auth/signup-form";

export default function Page() {
  return (
    <div className="flex">
      <div className="h-screen w-[45%] flex-shrink-0 scale-x-95 scale-y-[0.97] rounded-md bg-white"></div>
      <div className="my-auto flex w-full flex-col items-start px-32">
        <div>
          <h1 className="text-5xl font-semibold">Zentio</h1>
          <p className="text-neutral-500">Land Your Dream Tech Job With Ai</p>
        </div>
        <div className="w-full max-w-3xl pt-4">
          Let's create an Account
          <SignUpForm />
        </div>
      </div>
    </div>
  );
}
