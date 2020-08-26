package solace.coverage.paths;

import java.util.Arrays;
import java.util.LinkedList;
import java.util.List;
import java.util.Map;
import java.util.Queue;

public final class PathsTree<T> {
    private final String separator = "/";

    private Vertex<T> root = new Vertex<T>("", null);
    
    public <V extends Map<String, T>> PathsTree(V pathsList) {
        for (Map.Entry<String, T> key : pathsList.entrySet()) {
            this.addPath(key.getKey(), key.getValue());
        }        
    }

    private T resolve(Queue<String> parts, Vertex<T> current) {
        String key = parts.poll();
        Vertex<T> next = current.get(key);
        if (next == null) {
            List<Vertex<T>> variableChilgren = current.getVariableChildren();
            for (int i = 0; i < variableChilgren.size(); i++) {
                next = variableChilgren.get(i);
                if (next.matchesVariable(key)) {
                    if (parts.isEmpty())
                        return next.value;
                    T result = this.resolve(new LinkedList<String>(parts), next);
                    if (result != null)
                        return result;
                }
            }
            return null;
        }
        if (parts.isEmpty())
            return next.value;
        return resolve(parts, next);
    }

    public T resolve(String path) {
        return this.resolve(getPathPartsQueue(path), root);
    }

    private void addPath(String path, T value) {
        String[] parts = this.getPathParts(path);
        Vertex<T> current = root;
        Vertex<T> next;

        for (int i = 0; i < parts.length; i++) {
            String key = parts[i];
            next = current.get(key);
            if (next == null) 
                next = current.add(key, null);
            if (i == parts.length - 1) 
                next.value = value;
            current = next;
        }
    }

    private Queue<String> getPathPartsQueue(String path) {
        return new LinkedList<String>(Arrays.asList(getPathParts(path)));
    }

    private String[] getPathParts(String path) {
        if (path.startsWith(separator))
            return path.replaceFirst(separator, "").split(separator);
        return path.split(separator);
    }
}
