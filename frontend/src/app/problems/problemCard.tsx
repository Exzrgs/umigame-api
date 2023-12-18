import Image from 'next/image'
import './problemCard.css'

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
