module Main where

import Complex (complexMain)
import Control.Concurrent
import Control.Concurrent.Async
import Data.Foldable (traverse_)

main :: IO ()
main = complexMain

concurrentMain :: IO ()
concurrentMain = do
  filePaths <- getFilePaths
  mapConcurrently_ readAndPrintFileLines filePaths

syncMain :: IO ()
syncMain = do
  filePaths <- getFilePaths
  traverse_ readAndPrintFileLines filePaths

getFilePaths :: IO [FilePath]
getFilePaths = lines <$> getContents

readAndPrintFileLines :: FilePath -> IO ()
readAndPrintFileLines path = do
  file <- readFile path
  threadDelay 1000000
  let numLines = length . lines $ file
  putStrLn $ path ++ ": " ++ show numLines
