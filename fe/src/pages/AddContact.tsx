import { Input } from "@/components/ui/input";
import { useEffect, useState } from "react";
import { IoArrowBackSharp } from "react-icons/io5";
import { useDebounce } from "use-debounce";
import { MdOutlinePersonAdd } from "react-icons/md";
import { LoaderCircle } from "lucide-react";

function AddContactPage() {
  const [searchTerm, setSearchTerm] = useState("");
  const [search] = useDebounce(searchTerm, 1000);
  const [isError, setIsError] = useState(false);
  const [isLoad, setIsLoad] = useState(false);

  const handleSearchChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setIsLoad(true);
    setSearchTerm(e.target.value);
  };

  const searchUsername = async () => {
    setTimeout(() => {}, 1000);
    setIsError(false);
    if (search === "error") {
      setIsError(true);
    }
    console.log(search);
    setIsLoad(false);
  };

  useEffect(() => {
    if (search.length > 0) {
      searchUsername();
    }
    setIsLoad(false);
  }, [search]);

  return (
    <div>
      <div className="flex items-center">
        <button
          onClick={() => {
            window.location.replace("/contact");
          }}
          className="flex items-center p-2"
        >
          <IoArrowBackSharp size={26} />
        </button>
        <p className="text-2xl font-forta">Add Contact</p>
      </div>
      <form className="mt-4">
        <p className="mb-2 font-semibold">Username</p>
        <div className="relative w-full h-full">
          <Input
            placeholder="Masukan Username"
            className={`h-10`}
            onChange={handleSearchChange}
          />
          {isError && (
            <p className="mt-1 text-sm text-red-500 w-full text-left">
              Username yang anda masukan tidak valid
            </p>
          )}
          {isLoad && (
            <LoaderCircle className="animate-spin absolute text-zprimary right-4 top-2" />
          )}
        </div>
        <div className="mt-4">
          <button
            disabled={isLoad || isError}
            type="submit"
            className={`flex w-full gap-2 items-center justify-center cursor-pointer text-white font-medium rounded-md px-4 py-2 ease-in duration-300 ${
              isLoad || isError || search == ""
                ? "bg-gray-400"
                : "bg-zprimary hover:bg-hprimary"
            }`}
          >
            <p>Tambahkan</p>
            <MdOutlinePersonAdd size={26} />
          </button>
        </div>
      </form>
    </div>
  );
}

export default AddContactPage;
