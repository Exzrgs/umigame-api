import Image from 'next/image'
import Link from 'next/link'
import ProblemCard from './problemCard'
// ID               int       `json:"id"`
// 	Title            string    `json:"title"`
// 	ProblemStatement string    `json:"problem_statement"`
// 	Answer           string    `json:"answer"`
// 	CreatedAt        time.Time `json:"created_at"`
type Problem = {
    id: number,
    title: string,
    problemStatement: string,
    answer: string,
    createdAt: Date
}

type Chat = {
    id: number,
    title: string,
    problemStatement: string,
    answer: string,
    createdAt: Date
}

function truncateString(s: string, length: number) {
    if (s.length > length) {
        return s.substring(0, length - 3) + '...';
    } else {
        return s;
    }
}


const problems = [
    {
        id: 2,
        title: "aaaaa",
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
        id: 2,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    },
    {
        id: 3,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    },
    {
        id: 1,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    },
    {
        id: 2,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    },
    {
        id: 3,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    },
    {
        id: 1,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    },
    {
        id: 2,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    },
    {
        id: 3,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    },
    {
        id: 1,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
        answer: "ccccccc",
        createdAt: new Date('2000-04-01 10:00')
    },
    {
        id: 2,
        title: "aaaaa",
        problemStatement: "bbbbbbbb",
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

// export default function ProblemList() {
//     const list = problems.map((problem, i) => <li><Link legacyBehavior href={{ pathname: `/problem/${problem.id}` }}><ProblemCard key={i} id={problem.id} title={problem.title} problemStatement={problem.problemStatement} createdAt={problem.createdAt}></ProblemCard ></Link></li>)
//     return (
//         <>
//             <ul>
//                 {list}
//             </ul>
//         </>
//     )
// }
export default function ProblemList() {
    const cardWidth = 150; // カードの幅を設定
    const containerWidth = cardWidth * 3; // 親コンテナの幅を計算
    return (
        <div className="card-container">
            {problems.map((problem) => (
                <Link legacyBehavior href={{ pathname: `/problem/${problem.id}` }}>
                    <ProblemCard
                        key={problem.id}
                        id={problem.id}
                        title={problem.title}
                        problemStatement={truncateString(problem.problemStatement, 40)}
                        createdAt={problem.createdAt}
                    />
                </Link>
            ))}
        </div>
    );
}
