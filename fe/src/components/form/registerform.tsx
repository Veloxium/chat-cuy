import { useState } from "react";
import { delay, motion } from "framer-motion";
import { useForm } from "react-hook-form";
import { zodResolver } from "@hookform/resolvers/zod";
import loginSchema from "@/schema/loginSchema";
import { z } from "zod";
import { Input } from "@/components/ui/input";
import { LuEye, LuEyeOff } from "react-icons/lu";
import { FaGoogle } from "react-icons/fa";
import { Separator } from "@/components/ui/separator";
import registerSchema from "@/schema/registerSchema";

type RegisterFormInputs = z.infer<typeof registerSchema>;

function RegisterForm({
  setIsNewUser,
}: {
  setIsNewUser: (value: boolean) => void;
}) {
  const [isShowPassword, setIsShowPassword] = useState<boolean>(false);

  const container = {
    hidden: { opacity: 0 },
    show: {
      opacity: 1,
      transition: {
        delayChildren: 0.5,
        staggerChildren: 0.2,
      },
    },
    exit: {
      opacity: 0,
      transition: {
        delayChildren: 0.5,
        staggerChildren: 0.2,
      },
    },
  };

  const item = {
    hidden: { opacity: 0 },
    show: { opacity: 1 },
    exit: { opacity: 0 },
  };

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<RegisterFormInputs>({
    resolver: zodResolver(registerSchema),
  });

  const onSubmit = (data: RegisterFormInputs) => {
    console.log(data);
  };

  return (
    <motion.form
      onSubmit={handleSubmit(onSubmit)}
      variants={container}
      initial="hidden"
      animate="show"
      exit="exit"
      className="text-black flex flex-col gap-4 items-left w-full max-w-sm text-left"
    >
      <div>
        <motion.p variants={item} className="mb-2">
          Username
        </motion.p>
        <motion.div variants={item}>
          <Input
            placeholder="Masukan Username"
            className={`h-10 ${errors.username ? "border-red-500" : ""}`}
            {...register("username")}
          />
          {errors.username && (
            <p className="mt-1 text-red-500 w-full text-left">
              {errors.username.message}
            </p>
          )}
        </motion.div>
      </div>
      <div>
        <motion.p variants={item} className="mb-2">
          Email
        </motion.p>
        <motion.div variants={item}>
          <Input
            placeholder="Masukan Email"
            className={`h-10 ${errors.email ? "border-red-500" : ""}`}
            {...register("email")}
          />
          {errors.email && (
            <p className="mt-1 text-red-500 w-full text-left">
              {errors.email.message}
            </p>
          )}
        </motion.div>
      </div>
      <div>
        <motion.p variants={item} className="mb-2">
          Password
        </motion.p>
        <motion.div variants={item} className="relative place-content-center">
          <Input
            placeholder="Masukan Password"
            type={isShowPassword ? "text" : "password"}
            className={`h-10 ${errors.password ? "border-red-500" : ""}`}
            {...register("password")}
          />
          {errors.password && (
            <p className="mt-1 text-red-500 w-full text-left ">
              {errors.password.message}
            </p>
          )}
          {isShowPassword ? (
            <LuEye
              onClick={() => {
                setIsShowPassword(false);
              }}
              className="absolute z-10 top-0 right-2 text-slate-500 p-2 h-10 w-10"
            />
          ) : (
            <LuEyeOff
              onClick={() => {
                setIsShowPassword(true);
              }}
              size={26}
              className="absolute z-10 top-0 right-2 text-slate-500 p-2 h-10 w-10"
            />
          )}
        </motion.div>
      </div>
      <motion.div variants={item}>
        <button
          type="submit"
          className="flex w-full max-w-sm gap-2 items-center justify-center cursor-pointer"
        >
          <p className="text-lg border-2 border-zprimary bg-zprimary hover:bg-hprimary text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
            Daftar Cuy
          </p>
        </button>
      </motion.div>
      <motion.div
        variants={item}
        className="flex w-full gap-4 items-center justify-center cursor-pointer"
      >
        <button
          type="button"
          onClick={() => setIsNewUser(false)}
          className="flex w-full max-w-sm gap-2 group items-center justify-center"
        >
          <p className="bg-white border-2 border-zprimary group-hover:bg-hprimary text-zprimary group-hover:text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
            Sudah Punya Akun? Login
          </p>
        </button>
      </motion.div>
      <motion.div variants={item}>
        <Separator className="w-full max-w-sm" />
      </motion.div>
      <motion.div
        variants={item}
        className="flex w-full gap-4 items-center justify-center cursor-pointer"
      >
        <button
          className="flex w-full max-w-sm gap-2 group items-center justify-center"
          // onClick={() =>
          //   toast("Event has been created", {
          //     description: "Sunday, December 03, 2023 at 9:00 AM",
          //     action: {
          //       label: "Undo",
          //       onClick: () => console.log("Undo"),
          //     },
          //   })
          // }
        >
          <p className="bg-white border-2 border-zprimary group-hover:bg-hprimary text-zprimary group-hover:text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
            Dengan Google
          </p>
          <div className="bg-white border-2 border-zprimary group-hover:bg-hprimary rounded-md px-2 py-2 ease-in duration-300">
            <FaGoogle
              size={24}
              className="text-zprimary group-hover:text-white ease-in duration-300"
            />
          </div>
        </button>
      </motion.div>
    </motion.form>
  );
}

export default RegisterForm;
