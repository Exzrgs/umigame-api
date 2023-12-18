import Image from 'next/image'
import Link from 'next/link'
import { usePathname, useParams } from 'next/navigation'
import { useState } from "react"
import ChatCard from './chatCard';
type Problem = {
    id: number,
    title: string,
    problemStatement: string,
    answer: string,
    createdAt: Date
}

type Chat = {
    question: string,
    answer: string
}

const chatList = [
    {
        question: "これは誰でしょう？",
        answer: "どちらでもない"
    },
    {
        question: "これは誰でしょう？",
        answer: "はい"
    }
]
const problem =
{
    id: 1,
    title: "aaaaa",
    problemStatement: "bbbbbbbb",
    answer: "ccccccc",
    createdAt: new Date('2000-04-01 10:00')
}
export default function ChatList() {
    return (
        <>
            <li>
                {chatList.map((chat, i) => <ul><ChatCard key={i} question={chat.question} answer={chat.answer}></ChatCard></ul>)}
            </li>
        </>
    )
}
