import "./App.css";
import { Toaster } from "./components/ui/sonner";
import { AppRouter } from "./routes/AppRouter";
import { dynamicTitle } from "./utils/dynamicTitle";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

const queryClient = new QueryClient();

function App() {
  dynamicTitle();
  return (
    <QueryClientProvider client={queryClient}>
      <div className="App">
        <Toaster />
        <AppRouter />
      </div>
    </QueryClientProvider>
  );
}

export default App;
