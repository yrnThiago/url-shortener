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

function App() {

  const apiEndpoint = "http://localhost:3000/encurtaai"
  const headers = {
    'Content-Type': 'application/json'
  }

  const [shortUrl, setShortUrl] = useState("");

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
    try{
      const [response] = await Promise.all([
        axios.post(apiEndpoint, data, headers)
      ]);

      if (response.data.error) {
        throw new Error(response.data.error)
      }

      setShortUrl(response.data.message)
    } catch (error: any) {
      console.log(error)
    } finally {
      console.log("request sent")
    }
  }

  return (
    <div className="bg-neutral-950 flex flex-col items-center justify-center min-h-svh text-white">
        <Form {...form}>
          <form onSubmit={form.handleSubmit(handleShortUrl)} className="w-2/3 space-y-6">
            <FormField
              control={form.control}
              name="full_url"
              render={({ field }) => (
                <FormItem>
                  <FormLabel>URL</FormLabel>
                  <FormControl>
                    <Input placeholder="example.com" {...field} />
                  </FormControl>
                  <FormMessage />
                </FormItem>
              )}
            />
            <Button type="submit">Submit</Button>
          </form>
        </Form>
    </div>
  )
}

export default App;
