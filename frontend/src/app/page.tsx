'use client'
import ProblemList from './problems/page';
import VStack from './components/layouts/VStack';
import HStack from './components/layouts/HStack';
import Rectangle from './components/layouts/Rectangle';
import Color from './components/colors/Color';
import ScrollableContainer from './components/layouts/ScrollableContainer';
import ProblemCardList from './components/lists/ProblemCardList';
import Problem from './models/Problem/Problem';
import TabBar from './components/elements/tabBar';
import React, { useState, useEffect } from 'react'


export default function Home() {

  const [problems, setProblems] = useState<Problem[]>([])
  useEffect(() => {
    const fetchProblem = async () => {
      try {
        const response = await fetch('http://localhost:8080/problem/list?page=1')
        if (!response.ok) {
          throw new Error(`HTTP error! Status: ${response.status}`)
        }
        const data = await response.json();
        const convertedProblems = data.map((item: any) => ({
          id: item.ID,
          title: item.Title,
          author: item.Author,
          statement: item.Statement,
          isSolved: item.IsSolved,
          isLiked: item.IsLiked,
          createdAt: new Date(item.CreatedAt)
        }))
        setProblems(convertedProblems)
      } catch (error) {
        console.error('Failed to fetch problem:', error)
      }
    }

    fetchProblem();
  }, []) // 空の依存配列でマウント時に一度だけ実行

  return (
    <>
      <HStack style={{ height: '1080px', alignItems: 'flex-start' }}>
        <TabBar></TabBar>
        <Rectangle color={Color.Gray} width={160} height={1080}></Rectangle>
        <ScrollableContainer maxHeight={1080} maxWidth={1840} style={{ width: '100%' }}>
          <VStack>
            <Rectangle color={Color.Gray} width={1920} height={64}></Rectangle>
            <HStack gap={80}>
              <VStack gap={64}>
                <ProblemCardList problems={problems.slice(0, Math.ceil(problems.length / 2))}></ProblemCardList>
              </VStack>
              <VStack gap={64}>
                <ProblemCardList problems={problems.slice(Math.ceil(problems.length / 2), problems.length)}></ProblemCardList>
              </VStack>
            </HStack>
          </VStack >
        </ScrollableContainer >
      </HStack >
    </>
  )
}
