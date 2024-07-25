"use client"
import { useEffect, useRef } from "react"

export function DeleteAllCookies () {

  const callRef = useRef(false)

  useEffect(() => {
    if (callRef.current) return
    callRef.current = true
    fetch('/nextjs/api/cookies', { method: 'DELETE' })
  }, [])

  return null

}