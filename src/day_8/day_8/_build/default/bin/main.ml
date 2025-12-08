open Core

let rec find a reps =
  if not (phys_equal a reps.(a)) then reps.(a) <- find reps.(a) reps;
  reps.(a)
;;

let union a b reps = reps.(find a reps) <- reps.(find b reps)

let read_input filename =
  let ic = In_channel.create filename in
  let rec parser () =
    match In_channel.input_line ic with
    | None -> []
    | Some line ->
      let point = line |> String.split_on_chars ~on:[ ',' ] |> List.map ~f:Int.of_string in
      point :: parser ()
  in
  parser ()
;;

let part1 points =
  let reps = Array.init (List.length points) ~f:Fun.id in
  let edges =
    points
    |> List.concat_mapi ~f:(fun a_idx a ->
      points
      |> List.filter_mapi ~f:(fun b_idx b ->
        if b_idx <= a_idx
        then None
        else
          Some
            ( List.zip_exn a b |> List.fold_left ~init:0 ~f:(fun acc (x1, x2) -> acc + ((x1 - x2) * (x1 - x2)))
            , a_idx
            , b_idx
            , a
            , b )))
    |> List.sort ~compare:(fun (dst1, _, _, _, _) (dst2, _, _, _, _) -> Int.compare dst1 dst2)
  in
  List.take edges 1000 |> List.iter ~f:(fun (_, a_idx, b_idx, _, _) -> union a_idx b_idx reps);
  let components = Array.init (List.length points) ~f:(fun _ -> 0) in
  reps |> Array.iter ~f:(fun rep -> components.(find rep reps) <- 1 + components.(find rep reps));
  components |> Array.sort ~compare:(fun a b -> Int.compare b a);
  components.(0) * components.(1) * components.(2)
;;

let part2 points =
  let reps = Array.init (List.length points) ~f:Fun.id in
  let edges =
    points
    |> List.concat_mapi ~f:(fun a_idx a ->
      points
      |> List.filter_mapi ~f:(fun b_idx b ->
        if b_idx <= a_idx
        then None
        else
          Some
            ( List.zip_exn a b |> List.fold_left ~init:0 ~f:(fun acc (x1, x2) -> acc + ((x1 - x2) * (x1 - x2)))
            , a_idx
            , b_idx
            , a
            , b )))
    |> List.sort ~compare:(fun (dst1, _, _, _, _) (dst2, _, _, _, _) -> Int.compare dst1 dst2)
  in
  let rec krusteez edges =
    match edges with
    | [] -> failwith "urbad"
    | (_, a_idx, b_idx, a, b) :: tail ->
      let () = union a_idx b_idx reps in
      if reps |> Array.for_all ~f:(fun i -> phys_equal (find i reps) (find reps.(0) reps))
      then List.hd_exn a * List.hd_exn b
      else krusteez tail
  in
  edges |> krusteez
;;

"../input.txt" |> read_input |> part1 |> Printf.printf "%i\n";;
"../input.txt" |> read_input |> part2 |> Printf.printf "%i\n"
