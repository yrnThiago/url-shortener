import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { Button } from "@/components/ui/button"
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form"
import { Input } from "@/components/ui/input"
import axios from "axios"
import { useState } from "react"
import { Loader2 } from "lucide-react"

function App() {

  const apiEndpoint = "http://localhost:3000/"
  const headers = {
    'Content-Type': 'application/json'
  }

  const [shortUrl, setShortUrl] = useState("");
  const [isLoading, setIsLoading] = useState(false);

  const FormSchema = z.object({
    full_url: z.string().url({
      message: "URL needs to be valid.",
    }),
  });

  const form = useForm<z.infer<typeof FormSchema>>({
    resolver: zodResolver(FormSchema),
    defaultValues: {
      full_url: "",
    },
  });

  const handleShortUrl =  async(data: z.infer<typeof FormSchema>) => {
    setShortUrl("");
    setIsLoading(true);

    try{
      const [response] = await Promise.all([
        axios.post(apiEndpoint, data, headers)
      ]);

      if (response.data.error) {
        throw new Error(response.data.error)
      }

      setShortUrl(response.data.short_url)
    } catch (error: any) {
      console.log(error)
    } finally {
      setIsLoading(false);
    }
  }

  return (
    <div className="bg-neutral-950 flex flex-col items-center justify-center min-h-svh text-white">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(handleShortUrl)} className="w-1/2 space-y-6">
            <FormField
              control={form.control}
              name="full_url"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>URL</FormLabel>
                  <FormControl>
                    <div className="flex">
                      <Input placeholder="example.com" {...field} />
                      {!isLoading ? (
                        <Button className="bg-blue-500 cursor-pointer hover:bg-blue-700" type="submit">Encurtar</Button>
                      ) : (
                        <Button disabled className="bg-blue-700">Aguarde<Loader2 className="animate-spin" /></Button>
                      )}
                    </div>
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
          </form>
        </Form>

        {shortUrl &&  
          <Button variant="link">
            <a href={shortUrl}>
              {shortUrl}
            </a>
          </Button>
        }
    </div>
  )
}

export default App;
