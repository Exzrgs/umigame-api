'use client'
import Image from 'next/image'
import Link from 'next/link'
import { usePathname, useParams } from 'next/navigation'
import { useState } from "react"

export default function ChatCard(props: { question: string, answer: string }) {

    return (
        <>
            <p>Q. {props.question}</p>
            <p>A. {props.answer}</p>
        </>
    )
}
