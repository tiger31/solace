package solace.coverage.paths;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.HashMap;
import java.util.List;
import java.util.Map;
import java.util.regex.Pattern;
import java.util.stream.Collectors;

public final class Vertex<T> {
    private static final String isVariablePattern = ".*\\{.*\\}.*";
    private static final String variablePattern = "\\{.*\\}";
    private static final String variableReplacePatter = "(.+)";

    T value;
    private String key;
    private Map<String, Vertex<T>> tree = new HashMap<String, Vertex<T>>();
    private boolean isVariable;


    public Vertex(String key, T value) {
        this.key = key;
        this.value = value;
        this.isVariable = key.matches(isVariablePattern);
    }

    public Vertex<T> add(String key, T value) {
        Vertex<T> vertex = new Vertex<T>(key, value);
        this.tree.put(key, vertex);
        return vertex;
    }

    public Vertex<T> get(String key) {
        return this.tree.get(key);
    }
    

    public List<Vertex<T>> getVariableChildren() {
        return this.tree.values()
            .stream()
            .filter(v -> v.isVariable)
            .collect(Collectors.toList());
    }

    public boolean matchesVariable(String key) {
        return key.matches(getEscapedVariablePattern(this.key));
    }

    static String getEscapedVariablePattern(String key) {
        List<String> parts = new ArrayList<>(Arrays.asList(key.split(variablePattern)));
        parts.removeIf(part -> part.isEmpty());
        
        List<String> result = parts.stream()
            .map(part -> Pattern.quote(part))
            .collect(Collectors.toList());
        if (parts.size() == 2) { //If variable is in between plain text
            result.add(1, variableReplacePatter);
            return String.join("", result);
        } else if (parts.size() == 1) { //If at yhe start/end
            if (key.indexOf("{") == 0)
                result.add(0, variableReplacePatter);
            else
                result.add(variableReplacePatter);
            return String.join("", result);
        } else {
            return variableReplacePatter; 
        } 
    }

    public String getKey() {
        return this.key;
    }

    public T getValue() {
        return this.value;
    }

    public boolean isVariable() {
        return this.isVariable;
    }
}