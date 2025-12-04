open Core

let rec read_input () =
  match In_channel.input_line In_channel.stdin with
  | Some line -> Bytes.of_string line :: read_input ()
  | None -> []
;;

let part_2 =
  let dirs = [| -1, 0; -1, -1; -1, 1; 0, -1; 0, 1; 1, -1; 1, 0; 1, 1 |] in
  let input = read_input () |> Array.of_list in
  let num_rows = Array.length input in
  let num_cols = Bytes.length input.(0) in
  let prune_dirs row col =
    dirs
    |> Array.map ~f:(fun (i, j) ->
      if row + i < 0 || row + i >= num_rows || col + j < 0 || col + j >= num_cols
      then '.'
      else Bytes.get input.(row + i) (col + j))
    |> Array.count ~f:(fun (c : char) -> if phys_equal c '.' then false else true)
  in
  let searcher () =
    let rs = Sequence.init num_rows ~f:Fun.id in
    let cs = Sequence.init num_cols ~f:Fun.id in
    rs
    |> Sequence.fold ~init:0 ~f:(fun acc row ->
      let tmp =
        cs
        |> Sequence.filter_map ~f:(fun col ->
          if phys_equal (Bytes.get input.(row) col) '.'
          then None
          else (
            let num_adj = prune_dirs row col in
            if num_adj < 4
            then (
              let () = Bytes.set input.(row) col '.' in
              Some true)
            else Some false))
        |> Sequence.count ~f:Fun.id
      in
      tmp + acc)
  in
  let rec converger prev =
    let next = searcher () in
    if phys_equal prev next then prev else converger next + prev
  in
  converger 0
;;

let () = Printf.printf "%i\n" part_2
