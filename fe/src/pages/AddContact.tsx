import { Input } from "@/components/ui/input";
import axiosInstance from "@/utils/axiosInstance";
import { useEffect, useState } from "react";
import { IoArrowBackSharp } from "react-icons/io5";
import { useDebounce } from "use-debounce";
import { MdOutlinePersonAdd } from "react-icons/md";
import { LoaderCircle } from "lucide-react";
import { CiAt } from "react-icons/ci";
import { useNavigate } from "react-router-dom";
import { useQueryClient } from "@tanstack/react-query";

const API_URL = import.meta.env.VITE_API_URL;

function getAvatarUrl(avatar?: string) {
  if (!avatar || avatar === "") {
    return "https://api.dicebear.com/9.x/adventurer/svg?seed=Mason";
  }
  if (avatar.startsWith("http")) return avatar;
  return `${API_URL}${avatar}`;
}

interface SearchedUser {
  id: number;
  username: string;
  email: string;
  bio: string;
  avatar: string;
}

function AddContactPage() {
  const [searchTerm, setSearchTerm] = useState("");
  const [search] = useDebounce(searchTerm, 1000);
  const [isError, setIsError] = useState(false);
  const [errorMsg, setErrorMsg] = useState("");
  const [isLoad, setIsLoad] = useState(false);
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [foundUser, setFoundUser] = useState<SearchedUser | null>(null);
  const [successMsg, setSuccessMsg] = useState("");
  const navigate = useNavigate();
  const queryClient = useQueryClient();

  const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setIsLoad(true);
    setSearchTerm(e.target.value);
    setFoundUser(null);
    setIsError(false);
    setErrorMsg("");
    setSuccessMsg("");
  };

  const searchUsername = async () => {
    if (!search.trim()) return;
    setIsLoad(true);
    setIsError(false);
    setErrorMsg("");
    try {
      const res = await axiosInstance.get(`/api/users/search?username=${search}`);
      const users: SearchedUser[] = res.data.data;
      if (users && users.length > 0) {
        setFoundUser(users[0]);
      } else {
        setIsError(true);
        setErrorMsg("Username tidak ditemukan");
      }
    } catch {
      setIsError(true);
      setErrorMsg("Username tidak ditemukan");
    } finally {
      setIsLoad(false);
    }
  };

  useEffect(() => {
    if (search.length > 0) {
      searchUsername();
    } else {
      setIsLoad(false);
      setFoundUser(null);
    }
  }, [search]);

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    if (!foundUser) return;
    setIsSubmitting(true);
    setSuccessMsg("");
    setErrorMsg("");
    try {
      await axiosInstance.post("/api/contacts", {
        contact_username: foundUser.username,
      });
      setSuccessMsg("Kontak berhasil ditambahkan!");
      queryClient.invalidateQueries({ queryKey: ["getcontacts"] });
      setTimeout(() => navigate(-1), 1000);
    } catch (err: any) {
      const message = err.response?.data?.message || "Gagal menambahkan kontak";
      setErrorMsg(message);
      setIsError(true);
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div>
      <div className="flex items-center">
        <button
          onClick={() => {
            window.history.back();
          }}
          className="flex items-center p-2"
        >
          <IoArrowBackSharp size={26} />
        </button>
        <p className="text-2xl font-forta">Add Contact</p>
      </div>
      <form className="mt-4" onSubmit={handleSubmit}>
        <p className="mb-2 font-semibold">Username</p>
        <div className="relative w-full h-full">
          <Input
            placeholder="Masukan Username"
            className={`h-10 pl-10`}
            onChange={handleSearchChange}
          />
          {isError && (
            <p className="mt-1 text-sm text-red-500 w-full text-left">
              {errorMsg || "Username yang anda masukan tidak valid"}
            </p>
          )}
          {successMsg && (
            <p className="mt-1 text-sm text-green-600 w-full text-left">
              {successMsg}
            </p>
          )}
          {isLoad && (
            <LoaderCircle className="animate-spin absolute text-zprimary right-4 top-2" />
          )}
          <CiAt
            className={`absolute left-2 top-2 ${
              isError || search == "" ? "" : "text-zprimary"
            }`}
            size={24}
          />
        </div>
        {foundUser && (
          <div className="mt-3 p-3 rounded-md bg-white flex items-center gap-3">
            <div className="w-12 h-12 rounded-full border overflow-hidden bg-white">
              <img src={getAvatarUrl(foundUser.avatar)} alt="avatar" className="w-full h-full object-cover" />
            </div>
            <div>
              <p className="font-semibold">{foundUser.username}</p>
              {foundUser.bio && (
                <p className="text-sm text-gray-500">{foundUser.bio}</p>
              )}
            </div>
          </div>
        )}
        <div className="mt-4">
          <button
            disabled={isLoad || isError || !foundUser || isSubmitting}
            type="submit"
            className={`flex w-full gap-2 items-center justify-center cursor-pointer text-white font-medium rounded-md px-4 py-2 ease-in duration-300 ${
              isLoad || isError || !foundUser || isSubmitting
                ? "bg-gray-400 cursor-not-allowed"
                : "bg-zprimary hover:bg-hprimary"
            }`}
          >
            <p>{isSubmitting ? "Loading..." : "Tambahkan"}</p>
            <MdOutlinePersonAdd size={26} />
          </button>
        </div>
      </form>
    </div>
  );
}

export default AddContactPage;
