import Image from "next/image"

type EventImageProps = {
  src : string
  alt : string
}

export function EventImage ({ src, alt } : EventImageProps) {
  return (
    <Image
      src={src}
      alt={alt}
      width={277}
      height={277}
      priority
      className="rounded-2xl"
    />
  )
}