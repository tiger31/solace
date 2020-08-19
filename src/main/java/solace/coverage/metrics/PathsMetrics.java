package solace.coverage.metrics;

import java.util.Arrays;
import java.util.Map;

import com.sun.tools.javac.util.List;

import io.swagger.v3.oas.models.PathItem;
import io.swagger.v3.oas.models.Paths;
import solace.coverage.annotations.InnerMetric;
import solace.coverage.base.AbstractCollectionMetric;
import solace.coverage.base.IRequestProcessor;
import solace.coverage.base.IResponseProcessor;
import solace.coverage.paths.PathsTree;

public class PathsMetrics extends AbstractCollectionMetric<Paths, PathItem> implements IRequestProcessor, IResponseProcessor {
    private static final long serialVersionUID = -3879850642157283499L;
    @InnerMetric(name = "path")
    private PathItemMetric p;
    private PathsTree<PathItem> tree;

    public PathsMetrics(Paths paths) {
        super(paths);
        this.tree = new PathsTree<>(paths);
        this.p = new PathItemMetric(List.from(paths.values()).get(0));
        for (Map.Entry<String, PathItem> path : paths.entrySet()) {
            this.addMetric(new PathItemMetric(path.getValue()));
        }
    }
}