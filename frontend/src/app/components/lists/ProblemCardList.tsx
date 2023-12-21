import React from 'react';
import ProblemCard from '../elements/ProblemCard'
import Link from "next/link"
import Problem from '@/app/models/Problem/Problem';

interface ProblemListProps {
    problems: Problem[]
}
const ProblemCardList: React.FC<ProblemListProps> = ({ problems }) => {
    return (
        <div>
            {problems.map((problem) => (
                <Link href={`problem/${problem.id}`}>
                    <ProblemCard key={problem.id} problem={problem} />
                </Link>
            ))}
        </div>
    );
};

export default ProblemCardList;
