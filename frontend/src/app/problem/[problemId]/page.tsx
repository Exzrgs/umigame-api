'use client'
import Image from 'next/image'
import Link from 'next/link'
import { usePathname, useParams } from 'next/navigation'
import { useState } from "react"
import ChatList from './chatList';
import ProblemCard from './problemCard'
import './problemCard.css'
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

const problems = [
    {
        id: 1,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    },
    {
        id: 2,
        title: "cccca",
        problemStatement: `幼い息子のカメタは５ｍ崖上にいる父カメオを見上げていた。



        カメタ『パパ………登りたいよ………

        でもボクには無理みたいだ………』



        カメオは愛する息子を想えばこそ、

        心を鬼にしなければならない。



        カメタが自分の力で崖を這い上がってくるのを、

        カメオは険しい表情をして待った。



        だがカメタはカメオの期待に反して崖を登らず、

        その場を去ってしまった。



        背を向け去りゆく我が子を見て、

        カメオはカメタに感心し、見直した。



        一体なぜ！？`,
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    },
    {
        id: 3,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    }
]
export default function Problem() {

    const params = useParams()
    const problemId = params.problemId
    const id = Number(problemId)
    if (problems[id] == undefined) {
        return (

            <>
                <ul>
                    no problem found.
                </ul>
            </>
        )
    }
    return (

        <>
            <ProblemCard id={problems[id].id} title={problems[id].title} problemStatement={problems[id].problemStatement} createdAt={problems[id].createdAt}></ProblemCard>
            <ChatList></ChatList >
            <form>
                <label>
                    Question:
                    <input type="text" name="name" />
                </label>
                <input type="submit" value="Submit" />
            </form>
        </>
    )
}
