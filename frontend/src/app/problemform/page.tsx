'use client'
import React, { useState } from 'react';
import ProblemDetail from '../models/Problem/ProblemDetail';


export default function ProblemForm() {
    const [problem, setProblem] = useState<ProblemDetail>({
        id: 0,
        title: '',
        statement: '',
        answer: '',
        author: '',
        reference: '',
        referenceURL: '',
        createdAt: new Date()
    });

    const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setProblem({ ...problem, [e.target.name]: e.target.value });
    };

    const handleSubmit = (e: React.FormEvent) => {
        alert('submit！');
        e.preventDefault();
        fetch('http://localhost:8080/problem/add', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(problem)
        })
            .then(response => response.json())
            .then(data => console.log(data))
            .catch(error => console.error('Error:', error));
    };

    return (
        <form onSubmit={handleSubmit} >
            <input type="number" name="id" value={problem.id} onChange={handleChange} placeholder="ID" />
            <input type="text" name="title" value={problem.title} onChange={handleChange} placeholder="Title" />
            <input type="text" name="statement" value={problem.statement} onChange={handleChange} placeholder="Statement" />
            <input type="text" name="answer" value={problem.answer} onChange={handleChange} placeholder="Answer" />
            <input type="text" name="author" value={problem.author} onChange={handleChange} placeholder="Author" />
            <input type="text" name="reference" value={problem.reference} onChange={handleChange} placeholder="Reference" />
            <input type="text" name="referenceURL" value={problem.referenceURL} onChange={handleChange} placeholder="Reference URL" />
            <button type="submit" style={{ padding: '10px 20px', fontSize: '16px', backgroundColor: 'blue', color: 'white', border: 'none', borderRadius: '5px', cursor: 'pointer' }}>投稿</button>
        </form>
    )
};
