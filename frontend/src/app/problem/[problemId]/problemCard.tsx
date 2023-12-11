'use client'
import Image from 'next/image'
import Link from 'next/link'
import { usePathname, useParams } from 'next/navigation'
import { useState } from "react"
import ChatList from './chatList';
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

const problem =
{
    id: 1,
    title: "aaaaa",
    problemStatement: "bbbbbbbb",
    answer: "ccccccc",
    createdAt: new Date('2000-04-01 10:00')
}

export default function ProblemCard(props: { id: number, title: string, problemStatement: string, createdAt: Date }) {
    const japaneseDay = `${props.createdAt.getFullYear()}年${props.createdAt.getMonth() + 1}月${props.createdAt.getDate()}日`

    return (
        <div className="card">
            <div className="card-content">
                <div className="card-header">
                    <span className="card-id">{props.id}</span>
                    <h2 className="card-title">{props.title}</h2>
                </div>
                <div className="card-footer">
                    <p className="card-date">{japaneseDay}</p>
                    <p className="card-description">{props.problemStatement}</p>
                </div>
            </div>
        </div>
    )
}
