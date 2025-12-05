open Core

let part_1 () =
  let ic = In_channel.create "input.txt" in
  let rec parser () =
    match In_channel.input_line ic with
    | Some line ->
      if String.length line < 1
      then []
      else (
        let nums = String.split line ~on:'-' |> List.map ~f:Int.of_string |> Array.of_list in
        nums :: parser ())
    | None -> []
  in
  let input = parser () in
  let rec solve cnt =
    match In_channel.input_line ic with
    | Some line ->
      let num = Int.of_string line in
      List.count input ~f:(fun range -> range.(0) <= num && range.(1) >= num) |> min 1 |> ( + ) cnt |> solve
    | None ->
      In_channel.close ic;
      cnt
  in
  solve 0
;;

let () = Printf.printf "%i\n" (part_1 ())

let rec merge_intervals curr tail =
  match tail with
  | head :: tail ->
    if head.(0) >= curr.(0) && head.(0) <= curr.(1)
    then merge_intervals [| curr.(0); max head.(1) curr.(1) |] tail
    else curr :: merge_intervals head tail
  | [] -> [ curr ]
;;

let part_2 () =
  let ic = In_channel.create "input.txt" in
  let rec parser () =
    match In_channel.input_line ic with
    | Some line ->
      if String.length line < 1
      then (
        let () = In_channel.close ic in
        [])
      else (
        let nums = String.split line ~on:'-' |> List.map ~f:Int.of_string |> Array.of_list in
        nums :: parser ())
    | None -> []
  in
  let input = parser () in
  let sorted = List.sort input ~compare:(fun f g -> Int.compare f.(0) g.(0)) in
  (match sorted with
   | head :: tail -> merge_intervals head tail
   | [] -> [])
  |> List.fold ~init:0 ~f:(fun acc a -> acc + a.(1) - a.(0) + 1)
;;

let () = Printf.printf "%i\n" (part_2 ())
