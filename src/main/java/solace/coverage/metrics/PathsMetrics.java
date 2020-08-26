package solace.coverage.metrics;

import java.util.Map;

import io.swagger.v3.oas.models.PathItem;
import io.swagger.v3.oas.models.Paths;
import solace.coverage.base.AbstractMapMetric;
import solace.coverage.base.IMetric;
import solace.coverage.base.IRequestProcessor;
import solace.coverage.base.IResponseProcessor;
import solace.coverage.paths.PathsTree;

public class PathsMetrics extends AbstractMapMetric<Paths, PathItem> implements IRequestProcessor, IResponseProcessor {
    private static final long serialVersionUID = -3879850642157283499L;
    private PathsTree<IMetric<PathItem>> tree;

    public PathsMetrics(Paths paths) {
        super(paths);
        for (Map.Entry<String, PathItem> path : paths.entrySet()) {
            this.addMetric(path.getKey(), new PathItemMetric(path.getValue()));
        }
        this.tree = new PathsTree<>((Map<String, IMetric<PathItem>>)this);
    }
}