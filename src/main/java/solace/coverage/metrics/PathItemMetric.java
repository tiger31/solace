package solace.coverage.metrics;

import io.swagger.v3.oas.models.PathItem;
import solace.coverage.base.AbstractMetric;
import solace.coverage.base.IRequestProcessor;
import solace.coverage.base.IResponseProcessor;
import solace.coverage.base.MetricMetadata;

public class PathItemMetric extends AbstractMetric<PathItem> implements IRequestProcessor, IResponseProcessor {
    public PathItemMetric(PathItem entity) {
        super(entity);
    }

    private MetricMetadata metadata = new MetricMetadata();

    @Override
    public float getCoverage() {
        return 0;
    }

    @Override
    public void Reset() {
    }

    @Override
    public MetricMetadata getMetadata() {
        return this.metadata;
    }
    
}