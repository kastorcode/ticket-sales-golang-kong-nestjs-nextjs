import type { Metadata } from "next"
import { Inter } from "next/font/google"
import "./globals.css"
import { Footer } from "~/components/Footer"
import { Navbar } from "~/components/Navbar"

const inter = Inter({ subsets: ["latin"] })

export const metadata : Metadata = {
  title: "Ticket Sales System",
  description: "Ticket sales system with Golang, Kong, Nest.js and Next.js"
}

export default function RootLayout({ children } : Readonly<{children : React.ReactNode}>) {
  return (
    <html lang="en">
      <body className={`${inter.className} flex flex-col min-h-screen items-center bg-primary text-default`}>
        <div className="w-full max-w-[1256px] min-h-screen p-4 md:p-10">
          <Navbar />
          {children}
          <Footer />
        </div>
      </body>
    </html>
  )
}
