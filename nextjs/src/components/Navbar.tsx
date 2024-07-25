import Image from "next/image"
import Link from "next/link"

export function Navbar () {
  return (
    <div className="flex max-w-full items-center justify-items-stretch rounded-2xl bg-[rgb(192,192,192)] p-4 shadow-nav">
      <div className="flex grow items-center justify-center">
        <Link href="/"> 
          <Image
            src="nextjs/logo.svg"
            alt="tickets logo"
            width={149}
            height={32}
          />
        </Link>
      </div>
      <Link href={"/checkout"} className="min-h-6 min-w-6 grow-0 items-center">
        <Image
          src="nextjs/cart.svg"
          alt="cart icon"
          width={24}
          height={24}
        />
      </Link>
    </div>
  )
}
