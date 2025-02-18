import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import registerSchema from "@/schema/registerSchema";
import { zodResolver } from "@hookform/resolvers/zod";
import { motion } from "framer-motion";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { FaGoogle } from "react-icons/fa";
import { ImFacebook } from "react-icons/im";
import { LuEye, LuEyeOff } from "react-icons/lu";
import { z } from "zod";

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
      className="text-black flex flex-col gap-4 w-full max-w-3xl text-left"
    >
      <div className="flex flex-col md:grid md:grid-cols-2 w-full gap-4">
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
              <p className="mt-1 text-sm text-red-500 w-full text-left">
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
              <p className="mt-1 text-sm text-red-500 w-full text-left">
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
              <p className="mt-1 text-sm text-red-500 w-full text-left ">
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
        <div>
          <motion.p variants={item} className="mb-2">
            Konfirmasi Password
          </motion.p>
          <motion.div variants={item} className="relative place-content-center">
            <Input
              placeholder="Masukan Konfirmasi Password"
              type={isShowPassword ? "text" : "password"}
              className={`h-10 ${
                errors.konfirmasiPassword ? "border-red-500" : ""
              }`}
              {...register("konfirmasiPassword")}
            />
            {errors.konfirmasiPassword && (
              <p className="mt-1 text-sm text-red-500 w-full text-left ">
                {errors.konfirmasiPassword.message}
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
      </div>
      <motion.div variants={item}>
        <button
          type="submit"
          className="mt-2 flex w-full gap-2 items-center justify-center cursor-pointer"
        >
          <p className="border-2 border-zprimary bg-zprimary hover:bg-hprimary text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
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
          className="flex w-full gap-2 group items-center justify-center"
        >
          <p className="bg-white border-2 border-zprimary group-hover:bg-hprimary text-zprimary group-hover:text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
            Sudah Punya Akun? Login
          </p>
        </button>
      </motion.div>
      <motion.div variants={item}>
        <Separator className="w-full" />
      </motion.div>
      <motion.div
        variants={item}
        className="flex w-full gap-4 items-center justify-center cursor-pointer"
      >
        <button className="flex w-full gap-2 group items-center justify-center">
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
      <motion.div
        variants={item}
        className="flex w-full gap-4 items-center justify-center cursor-pointer"
      >
        <button className="flex w-full max-w-sm gap-2 group items-center justify-center">
          <p className="bg-white border-2 border-zprimary group-hover:bg-hprimary text-zprimary group-hover:text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
            Dengan Fesnuk
          </p>
          <div className="bg-white border-2 border-zprimary group-hover:bg-hprimary rounded-md px-2 py-2 ease-in duration-300">
            <ImFacebook
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
