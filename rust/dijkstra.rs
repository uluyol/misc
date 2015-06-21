#[derive(Copy,Clone)]
struct VertexID {
	id: usize,
	is_nil: bool,
}

impl std::fmt::Debug for VertexID {
	fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
		if self.is_nil {
			write!(f, "nil")
		} else {
			write!(f, "{}", self.id)
		}
	}
}

impl VertexID {
	fn new(id: usize) -> VertexID {
		VertexID{id: id, is_nil: false}
	}

	fn nil() -> VertexID {
		VertexID{id: 0, is_nil: true}
	}
}

type EdgeID = (VertexID, usize);

struct Graph<T> {
	pub vertices: Vec<T>,
	pub edges: Vec<Vec<(VertexID, f64)>>,
}

impl <T> Graph<T> {
	fn new() -> Graph<T> {
		Graph{
			vertices: Vec::new(),
			edges: Vec::new(),
		}
	}

	fn add_vertex(&mut self, v: T) -> VertexID {
		let vid = self.vertices.len();
		self.vertices.push(v);
		self.edges.push(Vec::new());
		VertexID::new(vid)
	}

	fn add_edge(&mut self, vid1: VertexID, vid2: VertexID, weight: f64) -> EdgeID {
		let eid = (vid1, self.edges.len());
		self.edges[vid1.id].push((vid2, weight));
		eid
	}
}

fn dijkstra<T>(graph: &Graph<T>, source: VertexID) -> Vec<(f64, VertexID)> {
	let mut dist_prev = Vec::with_capacity(graph.vertices.len());
	let mut queue = Vec::with_capacity(graph.vertices.len());
	for i in 0..graph.vertices.len() {
		let mut p = (::std::f64::INFINITY, VertexID::nil());
		if i == source.id {
			p = (0.0, VertexID::nil());
		}
		dist_prev.push(p);
		queue.push(VertexID::new(i));
	}

	while !queue.is_empty() {
		let mut minidx = 0;
		let mut mindist = dist_prev[queue[0].id].0;
		for i in 1..queue.len() {
			let vid = queue[i];
			if dist_prev[vid.id].0 < mindist {
				minidx = i;
				mindist = dist_prev[vid.id].0;
			}
		}
		let u = queue.swap_remove(minidx);
		for e in graph.edges[u.id].iter() {
			let alt = dist_prev[u.id].0 + e.1;
			if alt < dist_prev[e.0.id].0 {
				dist_prev[e.0.id] = (alt, u);
			}
		}
	}
	dist_prev
}

fn main() {
	let mut graph = Graph::new();
	let vid1 = graph.add_vertex(1);
	let vid2 = graph.add_vertex(1010);
	graph.add_edge(vid2, vid1, 1.4);
	let dist_prev = dijkstra(&graph, vid1);
	println!("{:?}", dist_prev);
	let dist_prev2 = dijkstra(&graph, vid2);
	println!("{:?}", dist_prev2);
}