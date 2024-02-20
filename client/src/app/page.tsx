/**
 * v0 by Vercel.
 * @see https://v0.dev/t/2WCxT8VzpCg
 * Documentation: https://v0.dev/docs#integrating-generated-code-into-your-nextjs-app
 */
import { Input } from "@/components/ui/input"
import { Button } from "@/components/ui/button"
import { TableHead, TableRow, TableHeader, TableCell, TableBody, Table } from "@/components/ui/table"

export default function Home() {
  return (
    <div className="bg-[#0f172a] min-h-screen flex flex-col items-center justify-center text-white">
      <section className="bg-[#1e293b] p-8 rounded-lg w-[800px] space-y-6">
        <header className="flex justify-between items-center border-b border-gray-600 pb-2">
          <h1 className="text-xl font-semibold">Round 1</h1>
          <span className="text-sm">2m left</span>
        </header>
        <div className="text-center">
          <h2 className="text-3xl font-bold mb-6">What is the most popular ice cream flavor?</h2>
          <form className="space-y-4">
            <Input placeholder="Name" />
            <Input placeholder="Answer" />
            <Button className="w-full bg-blue-600 hover:bg-blue-700">Submit</Button>
          </form>
        </div>
        <div className="pt-6">
          <h3 className="text-xl font-semibold mb-4">Last Round Results</h3>
            <Table>
              <TableHeader>
                <TableRow>
                  <TableHead className="w-[80px]">Round #</TableHead>
                  <TableHead>Question</TableHead>
                  <TableHead>Popular Answer 1</TableHead>
                  <TableHead>Popular Answer 2</TableHead>
                </TableRow>
              </TableHeader>
              <TableBody>
                <TableRow>
                  <TableCell className="font-medium">Round 1</TableCell>
                  <TableCell>What is your favorite color?</TableCell>
                  <TableCell>Alice, Bob: Blue</TableCell>
                  <TableCell>Charlie, Duke: Red</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell className="font-medium">Round 2</TableCell>
                  <TableCell>What is the best season?</TableCell>
                  <TableCell>Alice, Bob: Summer</TableCell>
                  <TableCell>Charlie, Duke: Winter</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell className="font-medium">Round 3</TableCell>
                  <TableCell>Which pet is the cutest?</TableCell>
                  <TableCell>Alice, Bob: Kittens</TableCell>
                  <TableCell>Charlie, Duke: Puppies</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell className="font-medium">Round 4</TableCell>
                  <TableCell>What is the most important invention?</TableCell>
                  <TableCell>Alice, Bob: Internet</TableCell>
                  <TableCell>Charlie, Duke: Electricity</TableCell>
                </TableRow>
                <TableRow>
                  <TableCell className="font-medium">Round 5</TableCell>
                  <TableCell>Which superhero would you want to be?</TableCell>
                  <TableCell>Alice, Bob: Superman</TableCell>
                  <TableCell>Charlie, Duke: Batman</TableCell>
                </TableRow>
              </TableBody>
            </Table>
        </div>
      </section>
    </div>
  )
}
