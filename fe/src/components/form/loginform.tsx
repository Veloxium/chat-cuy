import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import loginSchema from "@/schema/loginSchema";
import { ResponseInfo, User, useUserStore } from "@/store/userStore";
import { zodResolver } from "@hookform/resolvers/zod";
import { motion } from "framer-motion";
import { useState } from "react";
import { useForm } from "react-hook-form";
import { FaGoogle } from "react-icons/fa";
import { LuEye, LuEyeOff } from "react-icons/lu";
import { z } from "zod";
import { useToast } from "@/hooks/use-toast";
import axiosInstance from "@/utils/axiosInstance";
import { ImFacebook } from "react-icons/im";
import { signInWithPopup } from "firebase/auth";
import { auth, facebookProvider, googleProvider } from "@/lib/firebase";

type LoginFormInputs = z.infer<typeof loginSchema>;

function LoginForm({
   setIsNewUser,
}: {
   setIsNewUser: (value: boolean) => void;
}) {
   const [isShowPassword, setIsShowPassword] = useState<boolean>(false);
   const setUser = useUserStore((state) => state.setUser);
   const setAuthenticated = useUserStore((state) => state.setAuthenticated);

   const { toast } = useToast();

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
   } = useForm<LoginFormInputs>({
      resolver: zodResolver(loginSchema),
   });

   const onSubmit = async (data: LoginFormInputs) => {
      console.log(data);
      // document.location.href = "/chat";
      try {
         const res = await axiosInstance.post("/login", data);
         if (res.status === 200) {
            setUser(res.data);
            document.location.href = "/chat";
         } else {
            toast({
               title: "Login Gagal",
               description: "Email atau Password Salah",
               duration: 5000,
               variant: "destructive",
            });
         }
      } catch (error: unknown) {
         toast({
            title: "Login Gagal",
            description: "An unexpected error occurred" + error,
            duration: 5000,
            variant: "destructive",
         });
      }
   };

   const API_URL = "http://localhost:8080";

   const LoginWithFesnuk = async () => {
      try {
         const result = await signInWithPopup(auth, facebookProvider);
         const user = result.user;
         const accessToken = await user.getIdToken();
         const res = await fetch(`${API_URL}/withfacebook`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ accessToken }),
         });
         const resResult = (await res.json()) as ResponseInfo;
         const data = resResult.data;
         if (res.ok) {
            const user = {
               id: data.id,
               username: data.username,
               email: data.email,
               accessToken: data.accessToken,
               avatar: data.avatar,
               bio: data.bio,
               created_at: data.created_at,
            } satisfies User;

            setUser(user);
            setAuthenticated(true);

            // document.location.href = "/chat";
         }
      } catch (e: unknown) {
         console.log(e);
      }
   };

   const LoginWithGoogle = async () => {
      try {
         const result = await signInWithPopup(auth, googleProvider);
         const user = result.user;
         const accessToken = await user.getIdToken();
         const res = await fetch(`${API_URL}/withgoogle`, {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ accessToken }),
         });
         const resResult = (await res.json()) as ResponseInfo;
         const data = resResult.data;
         if (res.ok) {
            const user = {
               id: data.id,
               username: data.username,
               email: data.email,
               accessToken: data.accessToken,
               avatar: data.avatar,
               bio: data.bio,
               created_at: data.created_at,
            } satisfies User;

            setUser(user);
            setAuthenticated(true);

            document.location.reload();
         }
      } catch (e: unknown) {
         console.log(e);
      }
   };

   // const LoginWithGoogle = async () => {
   //    try {
   //       const result = await signInWithPopup(auth, googleProvider);
   //       const user = result.user;
   //       const accessToken = await user.getIdToken();
   //       console.log(accessToken);
   //       const res = await axiosInstance.post("/withgoogle", {
   //          body: { accessToken: accessToken },
   //       });
   //
   //       console.log(res);
   //       const data = res.data as ResponseInfo;
   //       if (data.success == true) {
   //          const user = {
   //             id: data.data.id,
   //             username: data.data.username,
   //             email: data.data.email,
   //             accessToken: data.data.accessToken,
   //             avatar: data.data.avatar,
   //             bio: data.data.bio,
   //             created_at: data.data.created_at,
   //          } satisfies User;
   //
   //          setUser(user);
   //          document.location.href = "/chat";
   //       }
   //    } catch (e: unknown) {
   //       console.log(e);
   //    }
   // };

   return (
      <motion.form
         onSubmit={handleSubmit(onSubmit)}
         variants={container}
         initial="hidden"
         animate="show"
         exit="exit"
         className="text-black flex flex-col gap-4 items-left w-full max-w-sm text-left"
      >
         <div className="flex flex-col md:grid md:grid-cols-1 w-full gap-4">
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
         </div>
         <motion.div variants={item}>
            <button
               type="submit"
               className="mt-2 flex w-full max-w-sm gap-2 items-center justify-center cursor-pointer"
            >
               <p className="border-2 border-zprimary bg-zprimary hover:bg-hprimary text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
                  Login Cuy
               </p>
            </button>
         </motion.div>
         <motion.div
            variants={item}
            className="flex w-full gap-4 items-center justify-center cursor-pointer"
         >
            <button
               type="button"
               onClick={() => setIsNewUser(true)}
               className="flex w-full max-w-sm gap-2 group items-center justify-center"
            >
               <p className="bg-white border-2 border-zprimary group-hover:bg-hprimary text-zprimary group-hover:text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
                  Belum Punya Akun? Daftar
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
               type="button"
               onClick={LoginWithGoogle}
               className="flex w-full max-w-sm gap-2 group items-center justify-center"
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
         <motion.div
            variants={item}
            className="flex w-full gap-4 items-center justify-center cursor-pointer"
         >
            <button
               type="button"
               onClick={LoginWithFesnuk}
               className="flex w-full max-w-sm gap-2 group items-center justify-center"
            >
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

export default LoginForm;
